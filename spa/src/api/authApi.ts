import axios from "axios";

const API_URL = "http://localhost:8000/";

export const loginUser = async (data: { email: string; password: string }) => {
    const response = await axios.post(`${API_URL}/sign-in`, data);
    return response.data;
};

export const registerUser = async (data: { email: string; password: string; login: string }) => {
    const response = await axios.post(`${API_URL}auth/sign-up`, data);
    return response.data;
};

export const codeUser = async (data: { email: string }) => {
    const response = await axios.post(`${API_URL}auth/code`, data);
    return response.data;
};

export const verificationUser = async (data: { code: string }) => {
    const response = await axios.post(`${API_URL}auth/verification`, data);
    return response.data;
};

export const passwordUser = async (data: { newPassword: string; repeatPassword: string }) => {
    const response = await axios.put(`${API_URL}auth/password`, data);
    return response.data;
};
