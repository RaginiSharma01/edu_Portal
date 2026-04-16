import { useState } from "react";
import { signupApi } from "../services/authApi";
import OTPModal from "../components/Auth/OTPModal";

const Signup = ({ role }: any) => {
  const [form, setForm] = useState<any>({});
  const [showOtp, setShowOtp] = useState(false);

  const handleChange = (e: any) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSignup = async () => {
    try {
      await signupApi({ ...form, role });
      setShowOtp(true);
    } catch {
      alert("Signup failed");
    }
  };

  return (
    <div>
      <h2>Signup ({role})</h2>

      <input name="firstName" onChange={handleChange} />
      <input name="email" onChange={handleChange} />
      <input name="password" onChange={handleChange} />

      <button onClick={handleSignup}>
        Create Account & Verify Email
      </button>

      {showOtp && (
        <OTPModal
          email={form.email}
          onClose={() => setShowOtp(false)}
        />
      )}
    </div>
  );
};

export default Signup;