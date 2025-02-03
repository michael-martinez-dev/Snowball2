import { Bill } from "../models/Bills";
import {
  CreateDebt,
  DeleteDebt,
  GetAllDebts,
  GetDebtByID,
  UpdateDebt,
} from "../../wailsjs/go/debt/DebtService";
import { debt } from "../../wailsjs/go/models";
import NewBill = debt.Debt;

export let getBills = async () => {
  const NewBills = await GetAllDebts();
  console.log(NewBills);
  if (NewBills === null) {
    return [];
  }
  const bills: Bill[] = NewBills.map((NewBill: any) => {
    return {
      id: NewBill.id,
      name: NewBill.name,
      type: NewBill.type,
      total: NewBill.total === "" ? 0.0 : parseFloat(NewBill.total),
      interest: NewBill.interest === "" ? 0.0 : parseFloat(NewBill.interest),
      monthlyMin:
        NewBill.monthlyMin === "" ? 0.0 : parseFloat(NewBill.monthlyMin),
      monthlyActual:
        NewBill.monthlyActual === "" ? 0.0 : parseFloat(NewBill.monthlyActual),
      dueDay: NewBill.dueDay,
    };
  });
  console.log(bills);
  return bills;
};

export let updateBill = async (bill: Bill) => {
  let TempBill = NewBill.createFrom(bill) as any;
  await UpdateDebt(TempBill);
};

export let deleteBill = async (id: number) => {
  await DeleteDebt(id);
};

export let createBill = async (bill: Bill) => {
  let TempBill = NewBill.createFrom(bill) as any;
  await CreateDebt(TempBill);
};
