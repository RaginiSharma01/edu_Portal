import axios from "axios"

const Api = "http://localhost:8090/api"

export const signupApi = (data: any)=> 
    axios.post(`${Api}/auth/signup` , data);


export const loginApi = (data: any) =>
  axios.post(`${Api}/auth/login`, data);

export const verifyOtpApi = (data: { email: string; otp: string }) =>
  axios.post(`${Api}/auth/verify-otp`, data);

export const resendOtpApi = (email: string) =>
  axios.post(`${Api}/auth/resend-otp`, { email });

//export const forgetPassword = ()