
import axios from 'axios';

export async function getTasks(token) {
    try {
        const response = await axios.get(`${process.env.EXPO_PUBLIC_BACKEND_URL}/tasks`, { headers: { 'Authorization': `Bearer ${token}`, "ngrok-skip-browser-warning": "69420" } });
        if (response.status === 200) {
            return { success: true, tasks: response.data.tasks }
        }
        return { success: false, message: response.data.message }
    } catch (error) {
        return { success: false, message: error.message }
    }
}
