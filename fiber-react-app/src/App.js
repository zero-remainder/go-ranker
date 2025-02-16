import React, { useState } from 'react';
import './App.css';

function App() {
  const [messages, setMessages] = useState([]); // Stores chat messages
  const [inputValue, setInputValue] = useState(''); // Stores the current input value

  // Function to handle sending a message
  const handleSendMessage = () => {
    if (inputValue.trim() === '') return; // Don't send empty messages

    // Add the user's message to the chat
    setMessages((prevMessages) => [
      ...prevMessages,
      { text: inputValue, sender: 'user' },
    ]);

    // Simulate a bot response (you can replace this with an API call to your Fiber backend)
    setTimeout(() => {
      setMessages((prevMessages) => [
        ...prevMessages,
        { text: 'This is a bot response.', sender: 'bot' },
      ]);
    }, 1000);

    // Clear the input field
    setInputValue('');
  };

  return (
      <div className="App">
        <header className="App-header">
          <h1>ChatGPT-like Chat</h1>
          <div className="chat-container">
            {/* Display chat messages */}
            {messages.map((message, index) => (
                <div
                    key={index}
                    className={`message ${message.sender === 'user' ? 'user-message' : 'bot-message'}`}
                >
                  {message.text}
                </div>
            ))}
          </div>
          {/* Input area for sending messages */}
          <div className="input-area">
            <input
                type="text"
                value={inputValue}
                onChange={(e) => setInputValue(e.target.value)}
                placeholder="Type a message..."
                onKeyPress={(e) => e.key === 'Enter' && handleSendMessage()}
            />
            <button onClick={handleSendMessage}>Send</button>
          </div>
        </header>
      </div>
  );
}

export default App;