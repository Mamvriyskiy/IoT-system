import axios from "axios";

const API_URL = "http://localhost:8000/";

export const loginUser = async (data: { email: string; password: string }) => {
    const response = await axios.post(`${API_URL}auth/sign-in`, data);
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

export const verificationUser = async (data: { code: string; token: string }) => {
    const response = await axios.post(`${API_URL}auth/verification`, data);
    return response.data;
};

export const passwordUser = async (data: { newPassword: string; token: string }) => {
    const response = await axios.put(`${API_URL}auth/password`, data);
    return response.data;
};

export const getListHome = async () => {
    const token = localStorage.getItem('jwt');
    console.log(token)
    console.log("=====")

    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.get(`${API_URL}api/homes/`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
    console.log(response.data)
    return response.data;
};

export const addHome = async (data: { name: string }) => {
    const token = localStorage.getItem('jwt');
    console.log(token)
    console.log("=====")

    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.post(`${API_URL}api/homes/`, data, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
    return response.data;
};

export const deleteHome = async ({ homeId }: { homeId: string }) => {
    const token = localStorage.getItem('jwt');
    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.delete(`${API_URL}api/homes/${homeId}`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
    return response.data;
};

export const addDevice = async (data: { name: string; homeID: string }) => {
    const token = localStorage.getItem('jwt');
    console.log(token)

    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.post(`${API_URL}api/homes/${data.homeID}/devices`, data, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token,
        }
    });

    return response.data;
};

export const getListDevice = async (homeID: string) => {
    const token = localStorage.getItem('jwt');
    console.log(token)

    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.get(`${API_URL}api/homes/${homeID}/devices`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
    console.log(response.data)
    return response.data;
};

export const addClient = async (data: { email: string; homeID: string }) => {
    const token = localStorage.getItem('jwt');
    console.log(token)

    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.post(`${API_URL}api/homes/${data.homeID}/accesses`, data, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token,
        }
    });

    return response.data;
};

export const getListClient = async (homeID: string) => {
    const token = localStorage.getItem('jwt');
    console.log(token)

    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.get(`${API_URL}api/homes/${homeID}/accesses`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
    console.log(response.data)
    return response.data;
};

export const deleteClient= async ({ accessID, homeId }: { accessID: string; homeId: string }) => {
    const token = localStorage.getItem('jwt');
    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.delete(`${API_URL}api/homes/${homeId}/accesses/${accessID}`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
    return response.data;
};

export const deleteDevice = async ({ deviceID, homeId }: { deviceID: string; homeId: string }) => {
    const token = localStorage.getItem('jwt');
    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.delete(`${API_URL}api/homes/${homeId}/devices/${deviceID}`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });

    return response.data;
};

export const startDevice= async ({ deviceID, homeId }: { deviceID: string; homeId: string }) => {
    const token = localStorage.getItem('jwt');
    if (!token) {
        throw new Error('Token not found');
    }

    console.log(token)
    const response = await axios.post(`${API_URL}api/homes/${homeId}/devices/${deviceID}/status`, null, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
    
    console.log('Response headers:', response.headers);

    return response.data;
};

export const getListHistory = async (homeID: string, deviceID: string) => {
    const token = localStorage.getItem('jwt');
    console.log(token)

    if (!token) {
        throw new Error('Token not found');
    }

    const response = await axios.get(`${API_URL}api/homes/${homeID}/devices/${deviceID}/history`, {
        headers: {
            'Content-Type': 'application/json',
            'Authorization': token
        }
    });
    console.log(response.data)
    return response.data;
};
