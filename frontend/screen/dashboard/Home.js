import React from "react";
import { Button, View, ScrollView, Text, StyleSheet, } from 'react-native';
import ProgressBar from "../../component/ProgressBar";
import Event from "../../component/Event";
import Icon from "../../component/Icon";

export default function HomeScreen () {
    return (
        <View style={styles.container}>
            <View style={{flex: 2}}>
                <ProgressBar title="My Progress" progress={100} subProgresses={[{ name: "Oops", value: 80 }, { name: "Oops", value: 60 }]} />
            </View>

            <View style={{ flex: 1.75, padding: 10 }}>
                <Text style={{fontSize: 20, fontWeight: "bold"}}>Quick Actions</Text>
                <ScrollView horizontal={true} showsHorizontalScrollIndicator={false}>
                    <Icon name={"Morning is here!"} key={1} execute={() => { }} image={"alarm-outline"} />
                    <Icon name={"Sync"} key={1} execute={() => { }} image={"git-pull-request"}/>
                    <Icon name={"Peace out"} key={1} execute={() => { }} image={"battery-dead"} />
                    <Icon name={"New Schedule"} key={1} execute={() => { }} image={"refresh"} />
                    <Icon name={"Holiday"} key={1} execute={() => { }} image={"airplane-outline"} />
                </ScrollView>
            </View>

            <View style={{ flex: 6}}>
                <Event heading={"Upcoming Events"} 
                    events={[{ name: "Fake Event", isCancelled: true, startTime: "2023-03-08T11:30:00", endTime: "2023-03-08T12:30:00" },
                    { name: "Badminton", isCancelled: false, startTime: "2023-03-08T18:30:00", endTime: "2023-03-08T20:30:00" }]}
                />
            </View>
        </View>
    )
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        padding: 10,
    },
});
