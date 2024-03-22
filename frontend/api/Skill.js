
import axios from 'axios';
import { BACKEND_URL } from "react-native-dotenv"

export async function getSkills(token) {
    try {
        const response = await axios.get(`${process.env.EXPO_PUBLIC_BACKEND_URL}/skills`, { headers: { 'Authorization': `Bearer ${token}`, "ngrok-skip-browser-warning": "69420" } });
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
        const response = await axios.post(`${process.env.EXPO_PUBLIC_BACKEND_URL}/skills`, skill, { headers: { 'Authorization': `Bearer ${token}`, "ngrok-skip-browser-warning": "69420" } });
        if (response.status === 200) {
            return { success: true, message: response.data.message }
        }

        return { success: false, message: response.data.message }
    } catch (error) {
        return { success: false, message: error.message }
    }
}

export async function getProposals(token, skillId) {
    try {
        const response = await axios.get(`${process.env.EXPO_PUBLIC_BACKEND_URL}/skills/${skillId}/proposals`, { headers: { 'Authorization': `Bearer ${token}`, "ngrok-skip-browser-warning": "69420" } });

        if (response.status === 200) {
            return { success: true, proposals: response.data.proposals }
        } else {
            return { success: false, proposals: [] }
        }
    } catch (error) {
        return { success: false, message: error.message }
    }
}

export async function cancelProposal(token, skillId, proposalId) {
    try {
        const response = await axios.delete(`${process.env.EXPO_PUBLIC_BACKEND_URL}/skills/${skillId}/cancel/${proposalId}`, { headers: { 'Authorization': `Bearer ${token}`, "ngrok-skip-browser-warning": "69420" } });
        if (response.status === 200) {
            return { success: true, message: response.data.message }
        } else {
            return { success: false }
        }
    } catch (error) {
        return { success: false, message: error.message }
    }
}