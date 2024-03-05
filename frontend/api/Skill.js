
import axios from 'axios';
import { BACKEND_URL } from "react-native-dotenv"

export async function getSkills(token) {
    try {
        const response = await axios.get(`${BACKEND_URL}/skills`, { headers: { 'Authorization': `Bearer ${token}` } });
        if (response.status === 200) {
            return { success: true, skills: response.data.skills }
        }
        return { success: false, message: response.data.message }
    } catch (error) {
        return { success: false, message: error.message }
    }
}

export async function createSkill(token, skill) {
    try {
        const response = await axios.post(`${BACKEND_URL}/skills`, skill, { headers: { 'Authorization': `Bearer ${token}` } });
        if (response.status === 200) {
            return { success: true, message: response.data.message }
        }

        return { success: false, message: response.data.message }
    } catch (error) {
        return { success: false, message: error.message }
    }
}