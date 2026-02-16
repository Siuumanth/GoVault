
import { useState } from 'react';
import { request } from '../api/client';
import { ENDPOINTS } from '../api/endpoints';
import { useAuth } from '../context/AuthContext';

export default function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const { login } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const data = await request(ENDPOINTS.AUTH.LOGIN, {
        method: 'POST',
        body: JSON.stringify({ email, password })
      });
      login({ username: data.username, email: data.email }, data.token);
    } catch (err) {
      alert("Login failed: " + err.message);
    }
  };

  return (
    <div className="min-h-screen bg-[#0d1117] flex items-center justify-center p-4">
      <form onSubmit={handleSubmit} className="w-full max-w-md bg-[#161b22] border border-[#30363d] p-8 rounded-2xl shadow-xl">
        <h1 className="text-3xl font-bold text-white mb-2 text-center">GoVault</h1>
        <p className="text-gray-400 mb-8 text-center text-sm">Secure Microservices Storage</p>
        
        <div className="space-y-4">
          <input 
            type="email" placeholder="Email" required
            className="w-full bg-[#0d1117] border border-[#30363d] p-3 rounded-lg text-white focus:border-blue-500 outline-none transition-all"
            onChange={(e) => setEmail(e.target.value)}
          />
          <input 
            type="password" placeholder="Password" required
            className="w-full bg-[#0d1117] border border-[#30363d] p-3 rounded-lg text-white focus:border-blue-500 outline-none transition-all"
            onChange={(e) => setPassword(e.target.value)}
          />
          <button className="w-full bg-blue-600 hover:bg-blue-500 text-white font-bold py-3 rounded-lg transition-colors mt-4">
            Sign In
          </button>
        </div>
      </form>
    </div>
  );
}