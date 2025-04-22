// internal/controllers/SEOController.go
package controllers

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SEOController struct{}

func NewSEOController() *SEOController {
	return &SEOController{}
}

type BasicSEOReport struct {
	URL             string    `json:"url"`
	Title           string    `json:"title"`
	TitleLength     int       `json:"title_length"`
	MetaDescription string    `json:"meta_description"`
	MetaDescLength  int       `json:"meta_desc_length"`
	H1Count         int       `json:"h1_count"`
	H2Count         int       `json:"h2_count"`
	ImageCount      int       `json:"image_count"`
	InternalLinks   int       `json:"internal_links"`
	ExternalLinks   int       `json:"external_links"`
	CanonicalURL    string    `json:"canonical_url"`
	SitemapExists   bool      `json:"sitemap_exists"`
	SitemapURLs     []string  `json:"sitemap_urls,omitempty"`
	AnalyzedAt      time.Time `json:"analyzed_at"`
}

func (c *SEOController) AnalyzeWebsite(ctx *fiber.Ctx) error {
	websiteURL := ctx.Query("url")
	if websiteURL == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "URL parameter is required",
		})
	}

	// Validate URL
	parsedURL, err := url.ParseRequestURI(websiteURL)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL format",
		})
	}

	// Fetch the webpage
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(websiteURL)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to fetch URL",
		})
	}
	defer resp.Body.Close()

	// Parse HTML
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to parse HTML",
		})
	}

	// Extract SEO data
	report := BasicSEOReport{
		URL:        websiteURL,
		AnalyzedAt: time.Now(),
	}

	// Title tag
	doc.Find("title").Each(func(i int, s *goquery.Selection) {
		report.Title = s.Text()
		report.TitleLength = len(s.Text())
	})

	// Meta description
	doc.Find("meta[name='description']").Each(func(i int, s *goquery.Selection) {
		if desc, exists := s.Attr("content"); exists {
			report.MetaDescription = desc
			report.MetaDescLength = len(desc)
		}
	})

	// Headings
	report.H1Count = doc.Find("h1").Length()
	report.H2Count = doc.Find("h2").Length()

	// Images
	report.ImageCount = doc.Find("img").Length()

	// Links
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		if href, exists := s.Attr("href"); exists {
			if strings.HasPrefix(href, "http") {
				if strings.Contains(href, parsedURL.Host) {
					report.InternalLinks++
				} else {
					report.ExternalLinks++
				}
			}
		}
	})

	// Canonical URL
	doc.Find("link[rel='canonical']").Each(func(i int, s *goquery.Selection) {
		if canonical, exists := s.Attr("href"); exists {
			report.CanonicalURL = canonical
		}
	})

	// Check for sitemap
	report.SitemapURLs, report.SitemapExists = c.checkSitemap(parsedURL)

	return ctx.JSON(report)
}

func (c *SEOController) checkSitemap(parsedURL *url.URL) ([]string, bool) {
	commonSitemapPaths := []string{
		"/sitemap.xml",
		"/sitemap_index.xml",
		"/sitemap.txt",
		"/sitemap.xml.gz",
		"/sitemap/sitemap.xml",
		"/sitemap",
	}

	var foundSitemaps []string
	client := &http.Client{Timeout: 5 * time.Second}

	for _, path := range commonSitemapPaths {
		sitemapURL := fmt.Sprintf("%s://%s%s", parsedURL.Scheme, parsedURL.Host, path)
		resp, err := client.Head(sitemapURL)
		if err != nil || resp.StatusCode != 200 {
			continue
		}

		// For XML sitemaps, verify the content
		if strings.HasSuffix(path, ".xml") {
			if !c.isValidXMLSitemap(sitemapURL, client) {
				continue
			}
		}

		foundSitemaps = append(foundSitemaps, sitemapURL)
	}

	return foundSitemaps, len(foundSitemaps) > 0
}

func (c *SEOController) isValidXMLSitemap(url string, client *http.Client) bool {
	resp, err := client.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	contentStr := string(content)
	return strings.Contains(contentStr, "<urlset") ||
		strings.Contains(contentStr, "<sitemapindex") ||
		strings.Contains(contentStr, "<?xml")
}
