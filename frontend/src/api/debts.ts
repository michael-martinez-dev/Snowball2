import {Bill} from "../models/Bills";
import {CreateDebtItem, DeleteDebtItem, RetrieveDebts, UpdateDebtItem} from "../../wailsjs/go/debt/Debt";
import {debt} from "../../wailsjs/go/models";
import GoBill = debt.GoBill;

export let getBills = async () => {
    const goBills = await RetrieveDebts();
    console.log(goBills);
    const bills: Bill[] = goBills.map((goBill: any) => {
        return {
            id: goBill.id,
            name: goBill.name,
            type: goBill.type,
            total: (goBill.total === "")? 0.00 : parseFloat(goBill.total),
            interest: (goBill.interest === "")? 0.00 : parseFloat(goBill.interest),
            monthlyMin: (goBill.monthlyMin === "")? 0.00 :  parseFloat(goBill.monthlyMin),
            monthlyActual: (goBill.monthlyActual === "")? 0.00 : parseFloat(goBill.monthlyActual),
            dueDay: goBill.dueDay,
        }
    });
    console.log(bills);
    return bills;
}

export let updateBill = async (bill: Bill) => {
    let goBill = (GoBill.createFrom(bill) as any);
    await UpdateDebtItem(
        goBill.id,
        goBill.dueDay,
        goBill.total?.toString(),
        goBill.monthlyMin?.toString(),
        goBill.monthlyActual?.toString(),
        goBill.interest?.toString(),
        goBill.name,
        goBill.type,
    );
}

export let deleteBill = async (id: number) => {
    await DeleteDebtItem(id);
}

export let createBill = async (bill: Bill) => {
    let goBill = (GoBill.createFrom(bill) as any);
    await CreateDebtItem(
        goBill.id,
        goBill.dueDay,
        goBill.total?.toString(),
        goBill.monthlyMin?.toString(),
        goBill.monthlyActual?.toString(),
        goBill.interest?.toString(),
        goBill.name,
        goBill.type,
    );
}
