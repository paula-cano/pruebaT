import axios from "axios";

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_URL || "http://localhost:8082",
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 5000, // Evita que una petición cuelgue indefinidamente
});

export const fetchData = async () => {
    try {
      const response = await apiClient.get("/data");
      return response.data;
    } catch (error) {
      if (error.response) {
        console.error(`Error ${error.response.status}:`, error.response.data);
      } else if (error.request) {
        console.error("No response received:", error.request);
      } else {
        console.error("Error setting up request:", error.message);
      }
      return [];
    }
  };

