<template>
<div class="debt-table">
    <v-data-table ref="table1" :headers="headers" :items="bills" :sort-by="sortBy" :group-by="groupBy" :search="search" class="elevation-10">
        <template v-slot:top>
            <v-toolbar flat>
                <v-toolbar-title><strong>Debt Snowball</strong></v-toolbar-title>
                <v-spacer></v-spacer>
                <v-checkbox v-model="group" label="Group by Type" color="primary" hide-details ></v-checkbox>
                <v-text-field v-model="search" label="Search" append-icon="mdi-magnify" single-line hide-details></v-text-field>
                <v-spacer></v-spacer>
                <v-btn variant="elevated" rounded @click="toggleTheme">
                  {{ oppositeTheme }}
                </v-btn>
                <v-spacer></v-spacer>
                <v-dialog v-model="dialog" max-width="500px">
                    <template v-slot:activator="{ props }">
                        <v-btn color="black" rounded elevation="5" class="mb-2" v-bind="props" icon="mdi-plus">

                        </v-btn>
                    </template>
                    <v-card>
                        <v-card-title>
                            <span class="text-h5">{{ formTitle }}</span>
                        </v-card-title>

                        <v-card-text>
                            <v-container fluid>
                                <v-row>
                                    <v-col cols="12" sm="6" md="4">
                                        <v-text-field v-model="form.name" label="Debt Name"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" md="4">
                                        <v-select v-model="form.type" :items="['Credit Card', 'Student Loan', 'Car Loan', 'Mortgage', 'Personal Loan', 'Other']" label="Debt Type"></v-select>
                                    </v-col>
                                    <v-col cols="12" sm="6" md="4">
                                        <v-text-field v-model.number="form.interest" label="Interest Rate" suffix="%"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" md="4">
                                        <v-text-field v-model.number="form.total" label="Total Amount" prefix="$"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" md="4">
                                        <v-text-field v-model.number="form.monthlyMin" label="Monthly Minimum" validation-value="1.00" prefix="$"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" md="4">
                                        <v-text-field v-model.number="form.monthlyActual" label="Monthly Actual" validation-value="1.00" prefix="$"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6" md="4">
                                        <v-text-field v-model.number="form.dueDay" label="Due Day"></v-text-field>
                                    </v-col>
                                </v-row>
                            </v-container>
                        </v-card-text>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="blue-darken-1" variant="text" @click="close">
                                Cancel
                            </v-btn>
                            <v-btn color="blue-darken-1" variant="text" @click="save">
                                Save
                            </v-btn>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
                <v-dialog v-model="dialogDelete" max-width="500px">
                    <v-card>
                        <v-card-title class="text-h5">Are you sure you want to delete this item?</v-card-title>
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="blue-darken-1" variant="text" @click="closeDelete">Cancel</v-btn>
                            <v-btn color="blue-darken-1" variant="text" @click="deleteItemConfirm">OK</v-btn>
                            <v-spacer></v-spacer>
                        </v-card-actions>
                    </v-card>
                </v-dialog>
            </v-toolbar>
        </template>
        <template v-slot:item.total="{ item }">
            <span v-if="item.raw.total > 0">
                {{ parseFloat(item.raw.total).toLocaleString('en-US', {style:"currency", currency:"USD"}) }}
            </span>
            <span v-else>🥳</span>
        </template>
        <template v-slot:item.interest="{ item }">
            <span v-if="item.raw.interest > 0">{{ item.raw.interest }}%</span>
            <span v-else>0</span>
        </template>
        <template v-slot:item.monthlyMin="{ item }">
            <span v-if="item.raw.monthlyMin > 0">{{ parseFloat(item.raw.monthlyMin).toLocaleString('en-US', {style:"currency", currency:"USD"}) }}</span>
            <span v-else>0</span>
        </template>
        <template v-slot:item.monthlyActual="{ item }">
            <span v-if="item.raw.monthlyActual > 0">{{ parseFloat(item.raw.monthlyActual).toLocaleString('en-US', {style:"currency", currency:"USD"}) }}</span>
            <span v-else>0</span>
        </template>
        <template v-slot:item.dueDay="{ item }">
            <span>{{ item.raw.dueDay }}</span>
            <!--        <span v-else>N/A</span>-->
        </template>
        <template v-slot:item.postPayment="{ item }">
            <span>{{ computePostPayment(item.raw.total,item.raw.interest, item.raw.monthlyActual).toLocaleString('en-US', {style:"currency", currency:"USD"}) }}</span>
        </template>
        <template v-slot:item.paymentsLeft="{ item }">
            <span>{{
          useActual?
            computePaymentsLeft(item.raw.total, item.raw.interest, item.raw.monthlyActual) :
            computePaymentsLeft(item.raw.total, item.raw.interest, item.raw.monthlyMin)
          }}</span>
        </template>
        <template v-slot:item.actions="{ item }">
            <v-icon small class="mr-2" @click="editItem(item.raw)">
                mdi-pencil
            </v-icon>
            <v-icon small @click="deleteItem(item.raw)">
                mdi-delete
            </v-icon>
        </template>
        <template v-slot:no-data>
            <v-btn color="primary" @click="reset">Reset</v-btn>
        </template>
    </v-data-table>
</div>
</template>

<script lang="ts">
import {
    defineComponent
} from "vue";
import {
    Bill
} from "../models/Bills";
import {
    useTheme
} from 'vuetify';

export default defineComponent({
    name: "DebtsTable",
    props: {
        bills: {
            type: Array as() => Bill[],
            required: true,
        },
        useActual: {
            type: Boolean,
            required: true,
        },
    },
    data() {
        return {
            search: "",
            group: true,
            sortBy: [{ key: 'total', order: 'asc' }],
            editedIndex: -1,
            form: {
                name: "",
                type: "",
                interest: 0.0,
                total: 0,
                monthlyMin: 0,
                monthlyActual: 0,
                dueDay: 0,
            } as Bill,
            defaultItem: {
                name: "",
                type: "",
                total: 0,
                monthlyMin: 0,
                monthlyActual: 0,
                dueDay: 0,
            } as Bill,
            headers: [
                { title: "Name", key: "name", align: "start", sortable: false },
                { title: "Type", key: "type" },
                { title: "Total Amount", key: "total" },
                { title: "Interest Rate", key: "interest" },
                { title: "Monthly Minimum",  key: "monthlyMin" },
                { title: "Monthly Actual", key: "monthlyActual" },
                { title: "Post Payment", key: "postPayment" },
                { title: "Payments Left", key: "paymentsLeft" },
                { title: "Due Day", key: "dueDay" },
                { title: "", key: "actions", sortable: false },
            ],
            dialog: false,
            dialogDelete: false,
        };
    },
    computed: {
        formTitle(): string {
            return this.editedIndex == -1 ? "New Bill" : "Edit Bill";
        },
        oppositeTheme(): string {
            const theme = useTheme()
            return theme.global.current.value.dark ? 'light theme' : 'dark theme'
        },
        groupBy(): any[] {
            if (this.group) {
                return [{ key: 'type', order: 'asc' }]
            }
            return []
        },
    },
    watch: {
        dialog(val: boolean) {
            val || this.close();
        },
        dialogDelete(val: boolean) {
            val || this.closeDelete();
        },
    },
    methods: {
        reset() {
            this.$emit("reset");
            this.editedIndex = -1;
            this.form = Object.assign({}, this.defaultItem);
            this.dialog = false;
            this.dialogDelete = false;
        },
        editItem(item: Bill) {
            this.editedIndex = item.id;
            this.form = Object.assign({}, item);
            this.dialog = true;
        },
        deleteItem(item: Bill) {
            this.editedIndex = item.id;
            this.form = Object.assign({}, item);
            this.dialogDelete = true;
        },
        deleteItemConfirm() {
            this.$emit("deleteDebt", this.form);
            this.closeDelete();
        },
        close() {
            this.dialog = false;
            this.form = Object.assign({}, this.defaultItem);
            this.editedIndex = -1;
        },
        closeDelete() {
            this.dialogDelete = false;
            this.$nextTick(() => {
                this.form = Object.assign({}, this.defaultItem);
                this.editedIndex = -1;
            });
        },
        save() {
            if (this.editedIndex > -1) {
                this.$emit("updateDebt", this.form);
            } else {
                this.$emit("addDebt", this.form);
            }
            this.close();
        },
        computeMonthlyInterest(total: number, interest: number): number {
            return total * (interest / 100) / 12;
        },
        computePostPayment(total: number, interest: number, monthlyMin: number): number {
            if (interest == undefined || interest == 0)
                return total - monthlyMin;
            return (total - monthlyMin + this.computeMonthlyInterest(total - monthlyMin, interest));
        },
        computePaymentsLeft(total: number, interest: number, monthlyMin: number): number {
            let paymentsLeft = 0;
            if (interest == undefined || interest == 0) {
                paymentsLeft = Math.ceil(total / monthlyMin)
            } else {
                paymentsLeft = Math.ceil((total - monthlyMin + this.computeMonthlyInterest(total - monthlyMin, interest)) / monthlyMin)
            }
            if (paymentsLeft < 0 || isNaN(paymentsLeft)) {
                paymentsLeft = 0
            }
            return paymentsLeft;
        },
    },
    setup() {
        const theme = useTheme()

        return {
            theme,
            toggleTheme: () => theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark',
        }
    },
});
</script>

<style>

</style>
