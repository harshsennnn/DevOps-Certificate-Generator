import axios from "axios";

const API_BASE = "http://localhost:8081";

export const generateCertificate = async (formData) => {
    const res = await axios.post(`${API_BASE}/generate-certificate`, formData, {
        responseType: "blob" // Important for file downloads
    });
    return res.data;
};
