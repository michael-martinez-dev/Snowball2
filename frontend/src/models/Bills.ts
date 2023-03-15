export interface Bill {
    id: number;
    name: string;
    type: string;
    total: number;
    interest?: number;
    monthlyMin: number;
    monthlyActual: number;
    dueDay: number;
}

export const mockBills: Bill[] = [
    {
        id: 1,
        name: "Rent",
        type: "Mortgage",
        interest: 3.05,
        total: 1000,
        monthlyMin: 100,
        monthlyActual: 100,
        dueDay: 1,
    },
    {
        id: 2,
        name: "Car Payment",
        type: "Loan",
        total: 500,
        interest: 0.05,
        monthlyMin: 50,
        monthlyActual: 50,
        dueDay: 15,
    },
    {
        id: 3,
        name: "Student Loan",
        type: "Loan",
        total: 10000,
        monthlyMin: 500,
        monthlyActual: 500,
        dueDay: 1,
    },
    {
        id: 4,
        name: "Credit Card",
        type: "Credit",
        total: 5000,
        monthlyMin: 200,
        monthlyActual: 200,
        dueDay: 15,
    },
    {
        id: 5,
        name: "Other",
        type: "Personal Loan",
        total: 1000,
        interest: 9.99,
        monthlyMin: 100,
        monthlyActual: 100,
        dueDay: 1,
    },
];