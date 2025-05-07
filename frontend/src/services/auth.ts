import { User } from "../types/User";

export async function signup(user: User) {
    const res = await fetch('http://localhost:8080/signup', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(user)
    });

    if (!res.ok) {
        throw new Error(await res.text());
    }

    return res.json();
}

