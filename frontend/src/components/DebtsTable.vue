<template>
    <div class="debt-table">
        <v-data-table
            ref="table1"
            :headers="headers"
            :items="debts"
            :sort-by="sortBy"
            :group-by="groupBy"
            :search="search"
            class="elevation-10"
        >
            <template v-slot:top>
                <v-toolbar flat>
                    <v-text-field
                        v-model="search"
                        label="Search"
                        append-icon="mdi-magnify"
                        single-line
                        hide-details
                        clearable
                    />
                    <v-spacer />
                    <v-checkbox
                        v-model="group"
                        label="Group by Type"
                        color="primary"
                        hide-details
                    ></v-checkbox>
                    <v-dialog v-model="dialog" max-width="500px">
                        <template v-slot:activator="{ props }">
                            <v-btn
                                color="black"
                                rounded
                                elevation="5"
                                class="mb-2"
                                v-bind="props"
                                icon="mdi-plus"
                            />
                        </template>
                        <v-card>
                            <v-card-title>
                                <span class="text-h5">{{ formTitle }}</span>
                            </v-card-title>
                            <v-card-text>
                                <v-container fluid>
                                    <v-row>
                                        <!-- Form Inputs Here -->
                                    </v-row>
                                </v-container>
                            </v-card-text>
                            <v-card-actions>
                                <v-btn
                                    color="blue-darken-1"
                                    variant="text"
                                    @click="closeDialog"
                                    >Cancel</v-btn
                                >
                                <v-btn
                                    color="blue-darken-1"
                                    variant="text"
                                    @click="save"
                                    >Save</v-btn
                                >
                            </v-card-actions>
                        </v-card>
                    </v-dialog>
                </v-toolbar>
            </template>
        </v-data-table>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from "vue";
import { useDebtsStore } from "../store/modules/debts";

export default defineComponent({
    name: "DebtsTable",
    props: {
        useActual: {
            type: Boolean,
            required: true,
        },
    },
    setup() {
        const debtsStore = useDebtsStore();
        const search = ref(""); // Search value
        const group = ref(false); // Grouping toggle
        const sortBy = ref([{ key: "total", order: "asc" }]); // Default sorting
        const dialog = ref(false); // Dialog visibility
        const headers = [
            { title: "Name", key: "name", align: "start", sortable: false }, // Disable sorting for "Name"
            { title: "Type", key: "type", sortable: true },
            { title: "Total Amount", key: "total", sortable: true },
            { title: "Interest Rate", key: "interest", sortable: true },
            { title: "Monthly Minimum", key: "monthlyMin", sortable: true },
            { title: "Monthly Actual", key: "monthlyActual", sortable: true },
            { title: "Post Payment", key: "postPayment", sortable: false },
            { title: "Payments Left", key: "paymentsLeft", sortable: false },
            { title: "Due Day", key: "dueDay", sortable: true },
            { title: "Actions", key: "actions", sortable: false }, // Actions column cannot be sorted
        ];
        const debts = computed(() => debtsStore.debts);

        const groupBy = computed(() => {
            return group.value ? [{ key: "type", order: "asc" }] : [];
        });

        const openDialog = () => {
            dialog.value = true;
        };

        const closeDialog = () => {
            dialog.value = false;
        };

        const save = () => {
            // Save logic
            closeDialog();
        };

        return {
            search,
            group,
            sortBy,
            dialog,
            debts,
            groupBy,
            openDialog,
            closeDialog,
            save,
            headers,
        };
    },
});
</script>
