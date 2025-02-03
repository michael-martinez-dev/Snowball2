<template>
    <v-sheet elevation="3" rounded="lg">
        <v-btn variant="elevated" @click="toggleTheme">
            {{ oppositeTheme }}
        </v-btn>
        <v-tabs
            v-model="tab"
            :items="tabs"
            align-tabs="center"
            color="grey"
            height="60"
            slider-color="#f78166"
        >
            <template v-slot:tab="{ item }">
                <v-tab
                    :text="item.text"
                    :value="item.value"
                    class="text-none"
                ></v-tab>
            </template>

            <template v-slot:item="{ item }">
                <v-tabs-window-item :value="item.value" class="pa-4">
                    <component :is="item.content" />
                </v-tabs-window-item>
            </template>
        </v-tabs>
    </v-sheet>
</template>

<script lang="ts">
import { defineComponent, markRaw } from "vue";
import NotesSection from "./NotesSection.vue";
import DebtsMain from "./DebtsMain.vue";
import HelpSection from "./HelpSection.vue";
import TotalsDisplay from "./TotalsDisplay.vue";
import { useTheme } from "vuetify";

const tabItems = [
    {
        text: "Debts",
        value: "debts",
        content: markRaw(DebtsMain),
    },
    {
        text: "Totals",
        value: "totals",
        content: markRaw(TotalsDisplay),
    },
    {
        text: "Notes",
        value: "notes",
        content: markRaw(NotesSection),
    },
    {
        text: "Help",
        value: "help",
        content: markRaw(HelpSection),
    },
];

export default defineComponent({
    name: "Home",
    data() {
        return {
            tab: "debts",
            tabs: tabItems,
        };
    },
    computed: {
        oppositeTheme(): string {
            const theme = useTheme();
            return theme.global.current.value.dark
                ? "light theme"
                : "dark theme";
        },
    },
    setup() {
        const theme = useTheme();

        return {
            theme,
            toggleTheme: () =>
                (theme.global.name.value = theme.global.current.value.dark
                    ? "light"
                    : "dark"),
        };
    },
});
</script>

<style scoped></style>
