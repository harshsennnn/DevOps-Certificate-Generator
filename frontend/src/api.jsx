import axios from "axios";

const API_BASE = "http://localhost:8081"; // your backend URL

export const fetchTemplates = () => {
  return axios.get(`${API_BASE}/templates`).then(res => res.data);
};

export const generateCertificate = (data) => {
  return axios.post(`${API_BASE}/generate-certificate`, data, {
    responseType: "blob" // for downloading PDF
  });
};
