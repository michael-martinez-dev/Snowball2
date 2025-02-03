<template>
    <div class="bottom-card">
        <v-card class="mx-auto" width="100%" height="100%" elevation="5">
            <v-card-title align="center">User Notes</v-card-title>
            <v-checkbox
                v-model="autoSave"
                label="Auto Save"
                @change="onUpdate"
            ></v-checkbox>
            <v-container fluid>
                <v-row cols="1" sm="6" md="4">
                    <v-textarea
                        v-model="userNotes"
                        label="My Notes"
                        variant="outlined"
                        @update:focused="onUpdate"
                        clearable
                    ></v-textarea>
                </v-row>
                <v-row cols="1" sm="6" md="4">
                    <v-btn
                        v-if="!autoSave"
                        color="blue-lighten-1"
                        variant="elevated"
                        @click="saveNotes"
                        location="center"
                    >
                        Save
                    </v-btn>
                </v-row>
            </v-container>
        </v-card>
    </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { UpdateNotes, ReadNotes } from "../../wailsjs/go/notes/NotesService";

export default defineComponent({
    name: "NotesSection",

    data() {
        return {
            notes: [
                "Pay off the highest interest rate debt first.",
                "Add debt using the plus button in the top right corner.",
                "Edit debt using the pencil icon.",
                "Delete debt using the trash can icon.",
                "You can sort by any field except Name.",
                "Add personal notes in the section above.",
                "Press the save button to save your notes.",
                "Press the x at the top right of the My Notes section to delete all text.",
            ] as string[],
            userNotesTitle: "DebtNotes",
            userNotes: "" as string,
            showNotes: false as boolean,
            autoSave: false as boolean,
        };
    },
    created() {
        this.getNotes();
    },
    methods: {
        saveNotes() {
            console.log("Saving notes...");
            UpdateNotes(this.userNotesTitle, this.userNotes).then(() => {
                this.getNotes();
            });
        },
        getNotes() {
            console.log("Getting notes...");
            ReadNotes(this.userNotesTitle).then((notes: string) => {
                this.userNotes = notes;
            });
        },
        onUpdate() {
            if (this.autoSave) {
                console.log("Auto saving notes...");
                this.saveNotes();
            }
        },
    },
});
</script>
