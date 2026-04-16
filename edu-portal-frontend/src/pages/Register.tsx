import { useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { signupApi } from "../services/authApi";
import OTPModal from "../components/Auth/OTPModal";
import "./Register.css";

const Register = () => {
  const { role } = useParams();
  const navigate = useNavigate();
  const [showOtp, setShowOtp] = useState(false);
  const [userType, setUserType] = useState<"student" | "teacher">("student");
  
  const [form, setForm] = useState({
    firstName: "",
    lastName: "",
    email: "",
    phone: "",
    age: "",
    dob: "",
    address: "",
    password: "",
    qualification: "",
    subjectsTeaching: "",
    fatherName: "",
    motherName: "",
    guardianName: "",
    occupation: "",
    height: "",
    weight: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleDobChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const dob = e.target.value;
    const birthDate = new Date(dob);
    const today = new Date();
    let age = today.getFullYear() - birthDate.getFullYear();
    const monthDiff = today.getMonth() - birthDate.getMonth();
    if (monthDiff < 0 || (monthDiff === 0 && today.getDate() < birthDate.getDate())) {
      age--;
    }
    setForm({ ...form, dob, age: age.toString() });
  };

  const validateForm = () => {
    if (!form.firstName) return alert("First name is required");
    if (!form.lastName) return alert("Last name is required");
    if (!form.email) return alert("Email is required");
    if (!form.phone) return alert("Phone number is required");
    if (!form.dob) return alert("Date of birth is required");
    if (!form.address) return alert("Address is required");
    if (!form.password) return alert("Password is required");
    
    if (userType === "student") {
      if (!form.fatherName) return alert("Father's name is required");
      if (!form.motherName) return alert("Mother's name is required");
      if (!form.occupation) return alert("Parent occupation is required");
      if (!form.height) return alert("Height is required");
      if (!form.weight) return alert("Weight is required");
    }
    
    if (userType === "teacher") {
      if (!form.qualification) return alert("Qualification is required");
      if (!form.subjectsTeaching) return alert("Subjects teaching is required");
    }
    
    return true;
  };

  const handleSubmit = async () => {
    if (!validateForm()) return;
    
    try {
      // Create clean payload
      let payload: any = {
        firstName: form.firstName,
        lastName: form.lastName,
        email: form.email,
        phone: form.phone,
        age: form.age ? parseInt(form.age) : undefined,
        dob: form.dob,
        address: form.address,
        password: form.password,
      };
      
      // Add student fields
      if (userType === "student") {
        payload.fatherName = form.fatherName;
        payload.motherName = form.motherName;
        if (form.guardianName) payload.guardianName = form.guardianName;
        payload.occupation = form.occupation;
        payload.height = form.height ? parseInt(form.height) : undefined;
        payload.weight = form.weight ? parseInt(form.weight) : undefined;
      }
      
      // Add teacher fields
      if (userType === "teacher") {
        payload.qualification = form.qualification;
        payload.subjectsTeaching = form.subjectsTeaching;
      }
      
      // Remove undefined values
      Object.keys(payload).forEach(key => {
        if (payload[key] === undefined || payload[key] === null || payload[key] === "") {
          delete payload[key];
        }
        if (typeof payload[key] === 'number' && isNaN(payload[key])) {
          delete payload[key];
        }
      });
      
      console.log("Sending payload:", payload);
      await signupApi(userType, payload);
      setShowOtp(true);
    } catch (error) {
      console.error("Registration error:", error);
      alert("Registration failed");
    }
  };

  return (
    <div className="register-container">
      <div className="register-card">
        <h2 className="register-title">Create account</h2>
        <p className="register-subtitle">
          Fill in your details — your email will be verified before you can login
        </p>

        <div className="role-tabs">
          <button
            className={`tab-btn ${userType === "teacher" ? "active" : ""}`}
            onClick={() => setUserType("teacher")}
          >
            Teacher
          </button>
          <button
            className={`tab-btn ${userType === "student" ? "active" : ""}`}
            onClick={() => setUserType("student")}
          >
            Student
          </button>
        </div>

        <div className="form-row">
          <input
            name="firstName"
            className="register-input"
            placeholder="First Name"
            value={form.firstName}
            onChange={handleChange}
          />
          
          <input
            name="lastName"
            className="register-input"
            placeholder="Last Name"
            value={form.lastName}
            onChange={handleChange}
          />
        </div>

        <input
          name="email"
          type="email"
          className="register-input"
          placeholder="Email"
          value={form.email}
          onChange={handleChange}
        />

        <input
          name="phone"
          className="register-input"
          placeholder="Phone"
          value={form.phone}
          onChange={handleChange}
        />

        <div className="form-row">
          <input
            name="age"
            className="register-input"
            placeholder="Age"
            value={form.age}
            readOnly
          />
          
          <input
            type="date"
            className="register-input"
            placeholder="Date of Birth"
            onChange={handleDobChange}
          />
        </div>

        <input
          name="address"
          className="register-input"
          placeholder="Full address"
          value={form.address}
          onChange={handleChange}
        />

        {userType === "student" && (
          <>
            <div className="form-row">
              <input
                name="fatherName"
                className="register-input"
                placeholder="Father's name"
                value={form.fatherName}
                onChange={handleChange}
              />
              
              <input
                name="motherName"
                className="register-input"
                placeholder="Mother's name"
                value={form.motherName}
                onChange={handleChange}
              />
            </div>

            <input
              name="guardianName"
              className="register-input"
              placeholder="Guardian name (if applicable)"
              value={form.guardianName}
              onChange={handleChange}
            />

            <div className="form-row">
              <input
                name="occupation"
                className="register-input"
                placeholder="Parent occupation"
                value={form.occupation}
                onChange={handleChange}
              />
              
              <input
                name="height"
                type="number"
                className="register-input"
                placeholder="Height (cm)"
                value={form.height}
                onChange={handleChange}
              />
              
              <input
                name="weight"
                type="number"
                className="register-input"
                placeholder="Weight (kg)"
                value={form.weight}
                onChange={handleChange}
              />
            </div>
          </>
        )}

        {userType === "teacher" && (
          <>
            <input
              name="qualification"
              className="register-input"
              placeholder="Qualification"
              value={form.qualification}
              onChange={handleChange}
            />
            
            <input
              name="subjectsTeaching"
              className="register-input"
              placeholder="Subjects teaching"
              value={form.subjectsTeaching}
              onChange={handleChange}
            />
          </>
        )}

        <input
          name="password"
          type="password"
          className="register-input"
          placeholder="Password"
          value={form.password}
          onChange={handleChange}
        />

        <button className="register-btn" onClick={handleSubmit}>
          Create Account & Verify Email
        </button>

        <p className="login-link">
          Already have an account?{" "}
          <a href="#" onClick={() => navigate(`/${userType}/login`)}>Sign in</a>
        </p>
      </div>

      {showOtp && <OTPModal email={form.email} onClose={() => setShowOtp(false)} />}
    </div>
  );
};

export default Register;