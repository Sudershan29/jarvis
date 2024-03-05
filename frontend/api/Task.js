
import axios from 'axios';
import { BACKEND_URL } from "react-native-dotenv"

export async function getTasks(token) {
    try {
        const response = await axios.get(`${BACKEND_URL}/tasks`, { headers: { 'Authorization': `Bearer ${token}` } });
        if (response.status === 200) {
            return { success: true, tasks: response.data.tasks }
        }
        return { success: false, message: response.data.message }
    } catch (error) {
        return { success: false, message: error.message }
    }
}
