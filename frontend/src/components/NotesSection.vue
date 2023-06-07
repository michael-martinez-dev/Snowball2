<template>
<div class="bottom-card">
    <v-card class="mx-auto" width="100%" height="100%" elevation="5">
        <v-checkbox v-model="autoSave" label="Auto Save" @change="onUpdate"></v-checkbox>
        <v-container fluid>
            <v-row cols="1" sm="6" md="4">
                <v-textarea v-model="userNotes" label="My Notes" variant="outlined" @update:focused="onUpdate" clearable></v-textarea>
            </v-row>
            <v-row cols="1" sm="6" md="4">
                <v-btn v-if="!autoSave" color="blue-lighten-1" variant="elevated" @click="saveNotes" location="center">
                    Save
                </v-btn>
            </v-row>
            <v-row cols="1" sm="6" md="4">
                <strong>Helpful Notes </strong>
                <v-btn color="blue-darken-1" variant="text" @click="showNotes = !showNotes">
                    <v-icon v-if="showNotes" left>
                        mdi-chevron-up
                    </v-icon>
                    <v-icon v-else left>
                        mdi-chevron-down
                    </v-icon>
                </v-btn>
            </v-row>
            <v-expand-transition>
                <div v-show="showNotes">
                    <v-card-text>
                        <ul>
                            <li v-for="note in notes" :key="note">
                                {{ note }}
                            </li>
                        </ul>
                    </v-card-text>
                </div>
            </v-expand-transition>
        </v-container>
    </v-card>
</div>
</template>

<script lang="ts">
import {
    defineComponent
} from "vue";
import {
    UpdateNotesBody,
    ReadNotes
} from "../../wailsjs/go/notes/Note"

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
            userNotes: "" as string,
            showNotes: false as boolean,
            autoSave: false as boolean,
        }
    },
    created() {
        this.getNotes();
    },
    methods: {
        saveNotes() {
            UpdateNotesBody(this.userNotes)
                .then(() => {
                    this.getNotes();
                })
        },
        getNotes() {
            ReadNotes().then((notes: string) => {
                this.userNotes = notes;
            });
        },
        onUpdate() {
            console.log("onUpdate");
            if (this.autoSave) {
                this.saveNotes();
                console.log("notes auto saved");
            }
        },
    },
})
</script>
