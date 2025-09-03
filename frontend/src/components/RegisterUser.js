import { useState } from "react";
import { registerUser } from "../api";

export default function RegisterUser({ onLogin }) {
  const [email, setEmail] = useState("");
  const [message, setMessage] = useState("");

  const handleLogin = async () => {
    if (!email) {
      setMessage("Email is required");
      return;
    }

    // Email format validation
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(email)) {
      setMessage("Invalid email format");
      return;
    }

    try {
      const res = await registerUser(email);
      if (res.user_id || res.message === "email exists") {
        setMessage("Successfully logged in!");
        onLogin({ email, user_id: res.user_id || null });
      } else {
        setMessage("Email not found. Please register.");
      }
    } catch (err) {
      setMessage("Error connecting to server");
    }
  };

  // Inline styles
  const containerStyle = {
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    height: "100vh",
    backgroundColor: "#f0f2f5",
  };

  const boxStyle = {
    backgroundColor: "#fff",
    padding: "40px 30px",
    borderRadius: "10px",
    boxShadow: "0 5px 15px rgba(0,0,0,0.2)",
    textAlign: "center",
    width: "300px",
  };

  const inputStyle = {
    width: "100%",
    padding: "10px",
    margin: "15px 0",
    borderRadius: "5px",
    border: "1px solid #ccc",
    fontSize: "16px",
  };

  const buttonStyle = {
    width: "100%",
    padding: "10px",
    backgroundColor: "#4caf50",
    color: "white",
    fontSize: "16px",
    border: "none",
    borderRadius: "5px",
    cursor: "pointer",
  };

  const messageStyle = {
    color: "red",
    marginTop: "10px",
    fontSize: "14px",
  };

  return (
    <div style={containerStyle}>
      <div style={boxStyle}>
        <h2>Login / Register</h2>
        <input
          type="email"
          placeholder="Enter your email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          style={inputStyle}
        />
        <button onClick={handleLogin} style={buttonStyle}>
          Login
        </button>
        {message && <p style={messageStyle}>{message}</p>}
      </div>
    </div>
  );
}
