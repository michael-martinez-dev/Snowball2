import { Bill } from "../models/Bills";
import {
  CreateDebtItem,
  DeleteDebtItem,
  RetrieveDebts,
  UpdateDebtItem,
} from "../../wailsjs/go/debt/Debt";
import { models } from "../../wailsjs/go/models";
import NewBill = models.NewBill;

export let getBills = async () => {
  const NewBills = await RetrieveDebts();
  console.log(NewBills);
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
  await UpdateDebtItem(
    TempBill.id,
    TempBill.dueDay,
    TempBill.total?.toString(),
    TempBill.monthlyMin?.toString(),
    TempBill.monthlyActual?.toString(),
    TempBill.interest?.toString(),
    TempBill.name,
    TempBill.type,
  );
};

export let deleteBill = async (id: number) => {
  await DeleteDebtItem(id);
};

export let createBill = async (bill: Bill) => {
  let TempBill = NewBill.createFrom(bill) as any;
  await CreateDebtItem(
    TempBill.id,
    TempBill.dueDay,
    TempBill.total?.toString(),
    TempBill.monthlyMin?.toString(),
    TempBill.monthlyActual?.toString(),
    TempBill.interest?.toString(),
    TempBill.name,
    TempBill.type,
  );
};
