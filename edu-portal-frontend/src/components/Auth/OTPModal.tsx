import { useState } from "react";
import { verifyOtpApi, resendOtpApi } from "../../services/authApi";

const OTPModal = ({ email, onClose }: any) => {
  const [otp, setOtp] = useState("");

  const handleVerify = async () => {
    try {
      await verifyOtpApi({ email, otp });
      alert("Verified!");
      onClose();
    } catch {
      alert("Invalid OTP");
    }
  };

  const handleResend = async () => {
    await resendOtpApi(email);
    alert("OTP resent");
  };

  return (
    <div className="modal-bg">
      <div className="modal">
        <h3>Enter OTP</h3>

        <input
          value={otp}
          onChange={(e) => setOtp(e.target.value)}
        />

        <button onClick={handleVerify}>Verify</button>
        <button onClick={handleResend}>Resend</button>
        <button onClick={onClose}>Close</button>
      </div>
    </div>
  );
};

export default OTPModal;