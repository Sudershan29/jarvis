
import React, {useContext, useState, useEffect} from "react";
import { Image, Button, View, Text, StyleSheet } from 'react-native';
import { GoogleCalendarConnectFn, getProfile, getCalendars } from "./../../api/Profile";
import { AuthContext } from "../../context/AuthContext";

export default function ProfileScreen ({navigation}) {
    const { userToken, logout } = useContext(AuthContext)
    const [user, setUser] = useState({})
    const [calendars, setCalendars] = useState([])

    useEffect(() => {
        const fetchData = async () => {
            const [userInfo, userCalendars] = await Promise.all([getProfile(userToken), getCalendars(userToken)]);
            setUser(userInfo);
            setCalendars(userCalendars);
        };

        fetchData();
    }, []);

    return (
        <View style={styles.container}>
            <Text style={styles.title}>My Profile</Text>

            <View style={styles.profilePictureContainer}>
                <Image
                    style={styles.profilePicture}
                    source={{uri: user?.profilePicture || 'https://via.placeholder.com/150'}}
                />
            </View>

            <View style={styles.userInfo}>
                <Text style={styles.userInfoText}>Name: {user?.name}</Text>
                <Text style={styles.userInfoText}>Email: {user?.email}</Text>
                <Text style={styles.userInfoText}>Role: {user?.role || 'N/A'}</Text>
                <Text style={styles.userInfoText}>Organization: {user?.organization || 'N/A'}</Text>
            </View>

            <Text style={styles.calendarTitle}>Calendars</Text>
            <View style={styles.calendarGrid}>
                {calendars.map((calendar, index) => (
                    <View key={index} style={styles.calendarItem}>
                        <Text style={styles.calendarText}>{calendar?.name}</Text>
                        <Text style={styles.calendarText}>{calendar?.type}</Text>
                        <Button
                            title="Remove"
                            onPress={() => {/* Function to remove calendar */}}
                            style={styles.removeButton}
                        />
                    </View>
                ))}
            </View>

            <Button
                title={calendars.length !== 0 ? "Replace Calendar" : "Connect to Calendar"}
                onPress={() => { window.location.href = GoogleCalendarConnectFn(userToken) }}
                style={styles.addButton}
            />
            
            <View style={{ marginBottom: 10 }}></View>
            
            <Button
                title="Logout"
                onPress={() => logout()}
                style={styles.logoutButton}
            />
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        alignItems: 'center',
        justifyContent: 'flex-start',
        paddingTop: 20,
        
    },
    title: {
        fontSize: 24,
        fontWeight: 'bold',
        marginBottom: 20,
    },
    profilePictureContainer: {
        marginBottom: 20,
    },
    profilePicture: {
        width: 150,
        height: 150,
        borderRadius: 75,
    },
    userInfo: {
        marginBottom: 20,
    },
    userInfoText: {
        fontSize: 18,
    },
    calendarTitle: {
        fontSize: 20,
        fontWeight: 'bold',
        marginBottom: 10,
    },
    calendarGrid: {
        marginBottom: 20,
    },
    calendarItem: {
        flexDirection: 'row',
        justifyContent: 'space-between',
        alignItems: 'center',
        padding: 10,
        marginVertical: 5,
        backgroundColor: '#fff',
        borderRadius: 10,
        width: 300,
    },
    calendarText: {
        fontSize: 16,
    },
    removeButton: {
        backgroundColor: '#ff6347',
        color: '#fff',
    },
    addButton: {
        marginBottom: 10,
    },
    logoutButton: {
        backgroundColor: '#ff6347',
        color: '#fff',
        margin: 10,
    },
});
