import { useState } from "react";
import { verifyOtpApi, resendOtpApi } from "../../services/authApi";
import { useNavigate } from "react-router-dom";
import "./OTPModal.css";

const OTPModal = ({ email, onClose }: any) => {
  const [otp, setOtp] = useState("");
  const [loading, setLoading] = useState(false);
  const [resendLoading, setResendLoading] = useState(false);
  const navigate = useNavigate();

  const handleVerify = async () => {
    if (!otp || otp.length !== 6) {
      alert("Please enter a valid 6-digit OTP");
      return;
    }

    setLoading(true);
    try {
      await verifyOtpApi({ email, otp });
      alert("Email verified successfully!");
      onClose();
      navigate("/dashboard");
    } catch {
      alert("Invalid OTP. Please try again.");
    } finally {
      setLoading(false);
    }
  };

  const handleResend = async () => {
    setResendLoading(true);
    try {
      await resendOtpApi(email);
      alert("OTP resent to your email!");
    } catch {
      alert("Failed to resend OTP. Please try again.");
    } finally {
      setResendLoading(false);
    }
  };

  return (
    <div className="modal-overlay" onClick={onClose}>
      <div className="modal-content" onClick={(e) => e.stopPropagation()}>
        <div className="modal-header">
          <h3>Verify Your Email</h3>
          <button className="modal-close" onClick={onClose}>×</button>
        </div>
        
        <div className="modal-body">
          <p className="modal-text">
            We've sent a verification code to<br />
            <strong>{email}</strong>
          </p>
          
          <input
            type="text"
            className="otp-input"
            placeholder="Enter 6-digit OTP"
            value={otp}
            onChange={(e) => setOtp(e.target.value.replace(/\D/g, '').slice(0, 6))}
            maxLength={6}
          />
          
          <div className="modal-buttons">
            <button 
              className="verify-button" 
              onClick={handleVerify}
              disabled={loading}
            >
              {loading ? "Verifying..." : "Verify Email"}
            </button>
            <button 
              className="resend-button" 
              onClick={handleResend}
              disabled={resendLoading}
            >
              {resendLoading ? "Sending..." : "Resend OTP"}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default OTPModal;