import { useState } from "react";
import { loginApi } from "../services/authApi";
import OTPModal from "../components/Auth/OTPModal";
import { useNavigate } from "react-router-dom";

const Login = ({ role }: any) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [showOtp, setShowOtp] = useState(false);

  const navigate = useNavigate();

  const handleLogin = async () => {
    try {
      const res = await loginApi({ email, password, role });

      if (!res.data.verified) {
        setShowOtp(true);
      } else {
        navigate("/dashboard");
      }
    } catch (err) {
      alert("Login failed");
    }
  };

  return (
    <div>
      <h2>Login</h2>

      <input value={email} onChange={(e) => setEmail(e.target.value)} />
      <input
        value={password}
        type="password"
        onChange={(e) => setPassword(e.target.value)}
      />

      <button onClick={handleLogin}>Sign In</button>

      {showOtp && (
        <OTPModal email={email} onClose={() => setShowOtp(false)} />
      )}
    </div>
  );
};

export default Login;