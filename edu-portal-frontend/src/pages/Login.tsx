import { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { loginApi } from "../services/authApi";
import OTPModal from "../components/Auth/OTPModal";
import "./Login.css";

const Login = () => {
  const { role } = useParams();
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showOtp, setShowOtp] = useState(false);

  const handleLogin = async () => {
    if (!email || !password) {
      alert("Please fill all fields");
      return;
    }

    try {
      const res = await loginApi({ email, password, role });
      if (!res.data?.verified) {
        setShowOtp(true);
      } else {
        navigate("/dashboard");
      }
    } catch {
      alert("Login failed");
    }
  };

  return (
    <div className="login-container">
      <div className="login-card">
        <h2 className="login-title">Welcome back</h2>
        <p className="login-subtitle">Sign in to your account</p>

        <input
          type="email"
          className="login-input"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        
        <input
          type="password"
          className="login-input"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
        />

        <button className="login-btn" onClick={handleLogin}>
          Sign In
        </button>

        {role !== "admin" && (
          <p className="register-link">
            Don't have an account?{" "}
            <a href="#" onClick={() => navigate(`/${role}/register`)}>Register</a>
          </p>
        )}
      </div>

      {showOtp && <OTPModal email={email} onClose={() => setShowOtp(false)} />}
    </div>
  );
};

export default Login;