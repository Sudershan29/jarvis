
import axios from 'axios';
import { BACKEND_URL } from "react-native-dotenv"

export async function userLogin(userEmail, password) {
    try {
        const response = await axios.post(`${BACKEND_URL}/login`, {
            email: userEmail,
            password: password
        });

        return response.data.token
    } catch (error) {
        return {};
    }
}