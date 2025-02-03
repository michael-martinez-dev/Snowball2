<template>
    <div class="debt-table">
        <v-switch v-model="useActual" label="Use Actual?" />

        <v-data-table
            ref="table1"
            :headers="headers"
            :items="debts"
            :sort-by="sortBy"
            :group-by="groupBy"
            :search="search"
            class="elevation-10"
        >
            <!-- Table Toolbar -->
            <template #top>
                <v-toolbar flat>
                    <!-- Search box -->
                    <v-text-field
                        v-model="search"
                        label="Search"
                        append-icon="mdi-magnify"
                        single-line
                        hide-details
                        clearable
                    />

                    <v-spacer />

                    <!-- Toggle Grouping -->
                    <v-checkbox
                        v-model="group"
                        label="Group by Type"
                        color="primary"
                        hide-details
                    ></v-checkbox>

                    <v-spacer />

                    <!-- Add New Debt -->
                    <v-btn
                        color="black"
                        rounded
                        elevation="5"
                        class="mb-2"
                        icon="mdi-plus"
                        @click="openDialogForNew"
                    />
                </v-toolbar>
            </template>

            <!-- Post Payment Column -->
            <template #item.postPayment="{ item }">
                {{ calculatePostPayment(item) }}
            </template>

            <!-- Payments Left Column -->
            <template #item.paymentsLeft="{ item }">
                {{ calculatePaymentsLeft(item) }}
            </template>

            <!-- Actions Column (Edit / Delete) -->
            <template #item.actions="{ item }">
                <v-btn
                    icon="mdi-pencil"
                    @click="openDialogForEdit(item)"
                    variant="text"
                />
                <v-btn
                    icon="mdi-delete"
                    color="red"
                    variant="text"
                    @click="deleteDebt(item.id)"
                />
            </template>
        </v-data-table>

        <!-- Dialog for Add/Edit -->
        <v-dialog v-model="dialog" max-width="500px">
            <v-card>
                <v-card-title>
                    <span class="text-h5">{{ formTitle }}</span>
                </v-card-title>

                <v-card-text>
                    <v-container fluid>
                        <v-row>
                            <!-- Name -->
                            <v-col cols="12" sm="6">
                                <v-text-field
                                    v-model="formData.name"
                                    label="Name"
                                    required
                                />
                            </v-col>

                            <!-- Type -->
                            <v-col cols="12" sm="6">
                                <v-select
                                    v-model="formData.type"
                                    :items="types"
                                    label="Type"
                                    required
                                />
                            </v-col>

                            <!-- Total Amount -->
                            <v-col cols="12" sm="6">
                                <v-text-field
                                    v-model="formData.total"
                                    prefix="$"
                                    label="Total Amount"
                                    type="number"
                                />
                            </v-col>

                            <!-- Interest Rate -->
                            <v-col cols="12" sm="6">
                                <v-text-field
                                    v-model="formData.interest"
                                    suffix="%"
                                    label="Interest Rate"
                                />
                            </v-col>
                        </v-row>

                        <v-row class="mt-2">
                            <!-- Monthly Min -->
                            <v-col cols="12" sm="4">
                                <v-text-field
                                    v-model="formData.monthlyMin"
                                    prefix="$"
                                    label="Monthly Min"
                                    type="number"
                                />
                            </v-col>

                            <!-- Button: Use Min -->
                            <v-col
                                cols="12"
                                sm="2"
                                class="d-flex align-center justify-center"
                            >
                                <v-btn
                                    color="primary"
                                    variant="outlined"
                                    size="small"
                                    @click="useMinAsActual"
                                >
                                    Use Min
                                </v-btn>
                            </v-col>

                            <!-- Monthly Actual -->
                            <v-col cols="12" sm="4">
                                <v-text-field
                                    v-model="formData.monthlyActual"
                                    prefix="$"
                                    label="Monthly Actual"
                                    type="number"
                                />
                            </v-col>
                        </v-row>

                        <!-- Due Day -->
                        <v-row class="mt-2">
                            <v-col cols="12" sm="6">
                                <v-text-field
                                    v-model="formData.dueDay"
                                    label="Due Day"
                                    type="number"
                                />
                            </v-col>
                        </v-row>
                    </v-container>
                </v-card-text>

                <v-card-actions>
                    <v-btn
                        color="blue-darken-1"
                        variant="text"
                        @click="closeDialog"
                    >
                        Cancel
                    </v-btn>
                    <v-btn color="blue-darken-1" variant="text" @click="save">
                        Save
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed, reactive } from "vue";
import { useDebtsStore } from "../store/modules/debts";

// A helper type for referencing each debt item
interface DebtItem {
    id?: number;
    name: string;
    type: string;
    total: string; // string but numeric content
    interest: string; // string but numeric content (annual %)
    monthlyMin: string; // string but numeric content
    monthlyActual: string; // string but numeric content
    dueDay: number;
}

export default defineComponent({
    name: "DebtsTable",
    props: {},
    setup() {
        const debtsStore = useDebtsStore();

        // Basic data
        const search = ref("");
        const group = ref(false);
        const sortBy = ref([{ key: "total", order: "asc" }]);
        const dialog = ref(false);
        const formTitle = ref("");
        const useActual = ref(true);

        // Allowed types for the "Type" field
        const types = [
            "CREDIT CARD",
            "STUDENT LOAN",
            "PERSONAL LOAN",
            "AUTO LOAN",
            "MORTGAGE",
            "OTHER",
        ];

        // Reactive form data for the dialog
        const formData = reactive<DebtItem>({
            id: 0,
            name: "",
            type: "",
            total: "",
            interest: "",
            monthlyMin: "",
            monthlyActual: "",
            dueDay: 0,
        });

        // Debts array from the store
        const debts = computed(() => debtsStore.debts);

        // Grouping logic
        const groupBy = computed(() => {
            return group.value ? [{ key: "type", order: "asc" }] : [];
        });

        // Table headers
        const headers = [
            { title: "Name", key: "name", align: "start", sortable: false },
            { title: "Type", key: "type", sortable: true },
            { title: "Total Amount", key: "total", sortable: true },
            { title: "Interest Rate", key: "interest", sortable: true },
            { title: "Monthly Min", key: "monthlyMin", sortable: true },
            {
                title: "Monthly Actual",
                key: "monthlyActual",
                sortable: true,
            },
            {
                title: "Post Payment",
                key: "postPayment",
                sortable: false,
            },
            {
                title: "Payments Left",
                key: "paymentsLeft",
                sortable: false,
            },
            { title: "Due Day", key: "dueDay", sortable: true },
            { title: "Actions", key: "actions", sortable: false },
        ];

        // Resets the form to blank
        const resetForm = () => {
            formData.id = 0;
            formData.name = "";
            formData.type = "";
            formData.total = "";
            formData.interest = "";
            formData.monthlyMin = "";
            formData.monthlyActual = "";
            formData.dueDay = 0;
        };

        // Open dialog for adding a new debt
        const openDialogForNew = () => {
            formTitle.value = "Add Debt";
            resetForm();
            dialog.value = true;
        };

        // Open dialog for editing an existing debt
        const openDialogForEdit = (debtItem: DebtItem) => {
            formTitle.value = "Edit Debt";
            formData.id = debtItem.id || 0;
            formData.name = debtItem.name;
            formData.type = debtItem.type;
            formData.total = debtItem.total;
            formData.interest = debtItem.interest;
            formData.monthlyMin = debtItem.monthlyMin;
            formData.monthlyActual = debtItem.monthlyActual;
            formData.dueDay = debtItem.dueDay;
            dialog.value = true;
        };

        // Close the dialog
        const closeDialog = () => {
            dialog.value = false;
        };

        // Sync monthlyActual with monthlyMin
        const useMinAsActual = () => {
            formData.monthlyActual = formData.monthlyMin;
        };

        // Save (create or update) a debt
        const save = () => {
            if (!formData.id) {
                // Add new debt
                debtsStore.addDebt({
                    name: formData.name,
                    type: formData.type,
                    total: formData.total.toString(),
                    interest: formData.interest.toString(),
                    monthlyMin: formData.monthlyMin.toString(),
                    monthlyActual: formData.monthlyActual.toString(),
                    dueDay: formData.dueDay,
                });
            } else {
                // Update existing
                debtsStore.updateDebt({
                    id: formData.id,
                    name: formData.name,
                    type: formData.type,
                    total: formData.total.toString(),
                    interest: formData.interest.toString(),
                    monthlyMin: formData.monthlyMin.toString(),
                    monthlyActual: formData.monthlyActual.toString(),
                    dueDay: formData.dueDay,
                });
            }
            dialog.value = false;
        };

        // Delete a debt by ID
        const deleteDebt = (id: number) => {
            debtsStore.removeDebt(id);
        };

        // ------------
        // Calculations
        // ------------

        /**
         * calculatePostPayment: principal + interest - payment
         * - interest is monthly interest, e.g. annualInterest / 12
         */
        const calculatePostPayment = (item: DebtItem): string => {
            // 1. parse numeric values
            let monthlyPayment;
            const principal = parseFloat(item.total) || 0;
            const monthlyInterestRate =
                (parseFloat(item.interest) || 0) / 100 / 12;

            if (useActual.value) {
                monthlyPayment = parseFloat(item.monthlyActual) || 0;
            } else {
                monthlyPayment = parseFloat(item.monthlyMin) || 0;
            }

            // 2. newPrincipal after 1 payment
            const newPrincipal =
                principal + principal * monthlyInterestRate - monthlyPayment;

            // 3. clamp at 0 for negative
            const result = Math.max(0, newPrincipal);

            // 4. return as currency string
            return `$${result.toFixed(2)}`;
        };

        /**
         * calculatePaymentsLeft: approximate # of months to fully pay off the debt
         * using the standard loan formula:
         *
         * n = log(A / (A - P * i)) / log(1 + i)
         *
         * where:
         *   P = principal
         *   i = monthly interest rate
         *   A = monthly payment
         *
         * if monthlyPayment <= monthlyInterest portion, then it's never fully paid => ∞
         */
        const calculatePaymentsLeft = (item: DebtItem): string => {
            let A;
            const principal = parseFloat(item.total) || 0;
            const i = (parseFloat(item.interest) || 0) / 100 / 12; // monthly interest

            if (useActual.value) {
                A = parseFloat(item.monthlyActual) || 0;
            } else {
                A = parseFloat(item.monthlyMin) || 0;
            }

            // Edge cases
            if (principal <= 0 || A <= 0) {
                return "N/A";
            }
            // If monthly payment can't cover monthly interest, no payoff
            if (A <= principal * i) {
                return "∞";
            }

            // No interest scenario
            if (i === 0) {
                const months = Math.ceil(principal / A);
                return months.toString();
            }

            // Standard formula
            // n = log(A / (A - P*i)) / log(1 + i)
            const numerator = Math.log(A / (A - principal * i));
            const denominator = Math.log(1 + i);
            let n = numerator / denominator;

            // Round up
            n = Math.ceil(n);
            // if n is extremely large, might want to clamp or show "∞"
            if (!isFinite(n)) {
                return "∞";
            }
            return n.toString();
        };

        return {
            // Data
            search,
            group,
            sortBy,
            dialog,
            formTitle,
            formData,
            types,
            useActual,

            // Computed
            debts,
            groupBy,

            // Table
            headers,

            // Methods
            openDialogForNew,
            openDialogForEdit,
            closeDialog,
            save,
            deleteDebt,
            resetForm,
            useMinAsActual,

            // Calculation methods (used in v-slot templates)
            calculatePostPayment,
            calculatePaymentsLeft,
        };
    },
});
</script>

<style scoped>
.debt-table {
    margin: 1rem;
}
</style>
