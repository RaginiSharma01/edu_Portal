import { useState } from "react";
import { useNavigate } from "react-router-dom";
import "./Landing.css";

const Landing = () => {
  const navigate = useNavigate();
  const [activeRole, setActiveRole] = useState("student");

  const roles = ["Student", "Teacher", "Admin"];

  const handleContinue = () => {
    if (activeRole === "admin") {
      navigate("/admin/login");
    } else {
      navigate(`/${activeRole.toLowerCase()}/register`);
    }
  };

  return (
    <div className="landing-container">
      <div className="landing-card">
        <h1 className="title">EduPortal</h1>
        <p className="subtitle">Select your role</p>

        <div className="role-slider">
          <div className="slider-bg">
            <div 
              className="slider-indicator"
              style={{
                transform: `translateX(${
                  activeRole === "student" ? "0%" : 
                  activeRole === "teacher" ? "100%" : "200%"
                })`
              }}
            />
            {roles.map((role) => (
              <button
                key={role}
                className={`role-option ${activeRole === role.toLowerCase() ? "active" : ""}`}
                onClick={() => setActiveRole(role.toLowerCase())}
              >
                {role}
              </button>
            ))}
          </div>
        </div>

        <button className="continue-btn" onClick={handleContinue}>
          Continue
        </button>
      </div>
    </div>
  );
};

export default Landing;