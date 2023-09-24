import React from "react";
import { Button, View, Text, StyleSheet, } from 'react-native';

export default function HomeScreen () {
    return (
        <View style={styles.container}>
            <Text>Home Screen</Text>
        </View>
    )
}

// export default function HomeScreen(props) {
//     const [request, response, promptAsync] = Google.useAuthRequest({
//         androidClientId: '40357062295-b776b2l5vsh5s4p5u7r892l3nuf05tq7.apps.googleusercontent.com',
//         iosClientId: '40357062295-tsd5pg68sgemb8vi357f7am2rb3djt10.apps.googleusercontent.com',
//         webClientId: '40357062295-6l7qufn2ki3fdgkcjmr85q61oa5d5qeb.apps.googleusercontent.com',
//     })
//     return (
//         <View style={styles.container}>
//             <Text>Home Screen</Text>
//             <Button
//                 title="Go to Profile"
//                 onPress={() => props.navigation.navigate('Profile')}
//             />
//             <Button title="Sign in" onPress={() => promptAsync()} />
//         </View>
//     )
// }

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
    },
});
