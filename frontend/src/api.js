const BASE_URL = "http://localhost:8080"; // Your Go backend

export const registerUser = async (email) => {
  const res = await fetch(`${BASE_URL}/users`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email }),
  });
  return res.json();
};

export const verifyNumber = async (userId, number) => {
  const res = await fetch(`${BASE_URL}/verify`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ user_id: userId, number }),
  });
  return res.json();
};

export async function getUserNumbers(userId) {
  if (!userId) return [];
  try {
    const res = await fetch(`http://localhost:8080/users/${userId}/numbers`);
    if (!res.ok) return [];
    const data = await res.json();
    console.log("Fetched numbers:", data); // debug
    return data; // data should be an array of { id, user_id, number, created_at }
  } catch (err) {
    console.error(err);
    return [];
  }
}


export const getAllUsers = async (page = 1, size = 10) => {
  const res = await fetch(`${BASE_URL}/users/all?page=${page}&size=${size}`);
  return res.json();
};
