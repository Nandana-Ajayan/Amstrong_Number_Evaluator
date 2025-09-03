import { useEffect, useState } from "react";
import { getUserNumbers } from "../api";

export default function UserNumbers({ user }) {
  const [numbers, setNumbers] = useState([]);

  useEffect(() => {
    const fetchNumbers = async () => {
      if (!user || !user.user_id) return;
      const nums = await getUserNumbers(user.user_id);
      setNumbers(nums);
    };
    fetchNumbers();
  }, [user]);

  if (!user) return null;

  return (
    <div style={{ marginTop: "20px" }}>
      <h3>Your Armstrong Numbers</h3>
      {numbers.length === 0 ? (
        <p>No numbers saved yet.</p>
      ) : (
        <ul>
          {numbers.map((num) => (
            <li key={num.id}>{num.number}</li>
          ))}
        </ul>
      )}
    </div>
  );
}
