import { BrowserRouter, Routes, Route } from "react-router-dom";
import Landing from "../pages/Landing";
import Login from "../pages/Login";
import Signup from "../pages/Signup";
//import Dashboard from "../pages/Dashboard";

const AppRoutes = () => {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Landing />} />
        <Route path="/student/onboarding" element={<Signup role="student" />} />
        <Route path="/teacher/onboarding" element={<Signup role="teacher" />} />
        <Route path="/admin/login" element={<Login role="admin" />} />
       {/* // <Route path="/dashboard" element={<Dashboard />} /> */}
      </Routes>
    </BrowserRouter>
  );
};

export default AppRoutes;