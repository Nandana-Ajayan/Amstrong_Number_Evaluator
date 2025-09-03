import { useEffect, useState } from "react";
import { getAllUsers } from "../api";

export default function AllUsers() {
  const [users, setUsers] = useState([]);

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const res = await getAllUsers();
        setUsers(res.users || []);
      } catch (err) {
        console.error(err);
      }
    };

    fetchUsers();
  }, []);

  return (
    <div>
      <h2>All Users & Their Numbers</h2>
      {users.length === 0 ? (
        <p>No users found.</p>
      ) : (
        <ul>
          {users.map((u) => (
            <li key={u.user_id}>
              <b>{u.email}</b>:{" "}
              {u.armstrong_numbers.map((n) => n.number).join(", ")}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
