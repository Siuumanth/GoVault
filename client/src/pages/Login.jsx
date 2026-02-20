import { useState } from 'react';
import { authApi } from '../api/auth';
import { useAuth } from '../context/AuthContext';

export default function Login() {
  const [isSignUp, setIsSignUp] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [username, setUsername] = useState('');
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const { login } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    setIsLoading(true);

    try {
      let data;
      if (isSignUp) {
        if (!username.trim()) {
          setError('Username is required');
          setIsLoading(false);
          return;
        }
        data = await authApi.signup(username.trim(), email.trim(), password);
        // After signup, automatically log in
        login({ username: data.username, email: data.email }, data.token);
      } else {
        data = await authApi.login(email.trim(), password);
        login({ username: data.username, email: data.email }, data.token);
      }
    } catch (err) {
      const errorMsg = err.message || 'Authentication failed';
      setError(errorMsg);
      // If login fails with "not found" or similar, suggest signup
      if (!isSignUp && (errorMsg.toLowerCase().includes('not found') || errorMsg.toLowerCase().includes('does not exist'))) {
        setTimeout(() => {
          if (confirm('Account not found. Would you like to sign up instead?')) {
            setIsSignUp(true);
          }
        }, 500);
      }
    } finally {
      setIsLoading(false);
    }
  };

  const switchMode = () => {
    setIsSignUp(!isSignUp);
    setError('');
    setEmail('');
    setPassword('');
    setUsername('');
  };

  return (
    <div className="min-h-screen bg-gv-dark flex items-center justify-center p-4">
      <div className="w-full max-w-md bg-[#161b22] border border-[#30363d] p-8 rounded-2xl shadow-xl">
        <h1 className="text-3xl font-bold text-white mb-2 text-center">GoVault</h1>
        <p className="text-gray-400 mb-6 text-center text-sm">Secure Microservices Storage</p>
        
        {/* Toggle between Login and Sign Up */}
        <div className="flex gap-2 mb-6 bg-gv-dark p-1 rounded-lg">
          <button
            type="button"
            onClick={() => isSignUp && switchMode()}
            className={`flex-1 py-2 px-4 rounded-md text-sm font-medium transition-all ${
              !isSignUp
                ? 'bg-blue-600 text-white'
                : 'text-gray-400 hover:text-white'
            }`}
          >
            Sign In
          </button>
          <button
            type="button"
            onClick={() => !isSignUp && switchMode()}
            className={`flex-1 py-2 px-4 rounded-md text-sm font-medium transition-all ${
              isSignUp
                ? 'bg-blue-600 text-white'
                : 'text-gray-400 hover:text-white'
            }`}
          >
            Sign Up
          </button>
        </div>

        <form onSubmit={handleSubmit} className="space-y-4">
          {isSignUp && (
            <input 
              type="text" 
              placeholder="Username" 
              required
              value={username}
              className="w-full bg-gv-dark border border-[#30363d] p-3 rounded-lg text-white focus:border-blue-500 outline-none transition-all"
              onChange={(e) => setUsername(e.target.value)}
            />
          )}
          <input 
            type="email" 
            placeholder="Email" 
            required
            value={email}
            className="w-full bg-gv-dark border border-[#30363d] p-3 rounded-lg text-white focus:border-blue-500 outline-none transition-all"
            onChange={(e) => setEmail(e.target.value)}
          />
          <input 
            type="password" 
            placeholder="Password" 
            required
            value={password}
            className="w-full bg-gv-dark border border-[#30363d] p-3 rounded-lg text-white focus:border-blue-500 outline-none transition-all"
            onChange={(e) => setPassword(e.target.value)}
          />
          
          {error && (
            <div className="p-3 bg-red-900/20 border border-red-500/30 rounded-lg text-red-400 text-sm">
              {error}
            </div>
          )}

          <button 
            type="submit"
            disabled={isLoading}
            className="w-full bg-blue-600 hover:bg-blue-500 disabled:opacity-50 disabled:cursor-not-allowed text-white font-bold py-3 rounded-lg transition-colors mt-4"
          >
            {isLoading ? 'Please wait...' : (isSignUp ? 'Sign Up' : 'Sign In')}
          </button>
        </form>
      </div>
    </div>
  );
}