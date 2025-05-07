import { useState } from 'react';
import { signup } from '../services/auth';

export default function SignupForm() {
    const [form, setForm] = useState({ username: '', email: '', password: ''});
    const [result, setResult] = useState<string | null>(null);

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setForm({...form, [e.target.name]: e.target.value })
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        const res = await signup(form);
        setResult(JSON.stringify(res));
    };

    return (
        <form onSubmit={handleSubmit}>
          <input name="username" onChange={handleChange} required />
          <input name="email" type="email" onChange={handleChange} required />
          <input name="password" type="password" onChange={handleChange} required />
          <button type="submit">Sign Up</button>
          <div>{result}</div>
        </form>
      );
}


