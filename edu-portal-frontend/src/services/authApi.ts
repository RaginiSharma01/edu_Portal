import axios from "axios";

const BASE_API = "http://localhost:8090/api/smp/v1/onboarding";
const OTP_API = "http://localhost:9005";

export const signupApi = (role: string, data: any) =>
  axios.post(`${BASE_API}/signup/${role}`, data);

export const loginApi = (data: any) =>
  axios.post(`${BASE_API}/login`, data);

export const verifyOtpApi = (data: { email: string; otp: string }) =>
  axios.post(`${BASE_API}/verify-otp`, data);

export const resendOtpApi = (email: string) =>
  axios.post(`${OTP_API}/resend-otp`, { email });