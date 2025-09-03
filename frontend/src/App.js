import { useState } from "react";
import RegisterUser from "./components/RegisterUser";
import VerifyNumber from "./components/VerifyNumber";
import UserNumbers from "./components/UserNumbers";
import AllUsers from "./components/AllUsers";
import "./App.css";

function App() {
  const [user, setUser] = useState(null); // null = not signed in
  const [activeTab, setActiveTab] = useState("verify"); // default tab

  // If user not signed in, show only login/register
  if (!user) {
    return <RegisterUser onLogin={setUser} />;
  }

  // After login/registration
  return (
    <div className="App">
      <h1>Welcome, {user.email}</h1>

      <nav>
        <button onClick={() => setActiveTab("verify")}>Verify Number</button>
        <button onClick={() => setActiveTab("userNumbers")}>My Numbers</button>
        <button onClick={() => setActiveTab("allUsers")}>All Users</button>
        <button onClick={() => setUser(null)}>Logout</button>
      </nav>

      <div className="tabContent">
        {activeTab === "verify" && <VerifyNumber user={user} />}
        {activeTab === "userNumbers" && <UserNumbers user={user} />}
        {activeTab === "allUsers" && <AllUsers />}
      </div>
    </div>
  );
}

export default App;
