import { useState } from "react";
import { verifyNumber } from "../api";

export default function VerifyNumber({ user }) {
  const [number, setNumber] = useState("");
  const [result, setResult] = useState("");

  const handleVerify = async () => {
    if (!number) return setResult("Number is required");

    try {
      const res = await verifyNumber(user.user_id, parseInt(number));
      if (res.is_armstrong) {
        setResult("✅ Armstrong! Saved to your numbers.");
      } else {
        setResult("❌ Not Armstrong. Not saved.");
      }
    } catch (err) {
      setResult("Error verifying number");
      console.error(err);
    }
  };

  return (
    <div>
      <h2>Verify Number</h2>
      <input
        type="number"
        value={number}
        onChange={(e) => setNumber(e.target.value)}
        placeholder="Enter number"
      />
      <button onClick={handleVerify}>Verify</button>
      {result && <p>{result}</p>}
    </div>
  );
}
