import { useNavigate } from "react-router-dom";

const Landing = () => {
  const navigate = useNavigate();

  return (
    <div>
      <h1>EduPortal</h1>

      <button onClick={() => navigate("/student/onboarding")}>
        Student
      </button>

      <button onClick={() => navigate("/teacher/onboarding")}>
        Teacher
      </button>

      <button onClick={() => navigate("/admin/login")}>
        Admin
      </button>
    </div>
  );
};

export default Landing;