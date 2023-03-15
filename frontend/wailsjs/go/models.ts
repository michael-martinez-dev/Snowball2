export namespace debt {
	
	export class GoBill {
	    id: number;
	    dueDay: number;
	    name: string;
	    type: string;
	    total: string;
	    monthlyMin: string;
	    monthlyActual: string;
	    interest: string;
	
	    static createFrom(source: any = {}) {
	        return new GoBill(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.dueDay = source["dueDay"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.total = source["total"];
	        this.monthlyMin = source["monthlyMin"];
	        this.monthlyActual = source["monthlyActual"];
	        this.interest = source["interest"];
	    }
	}

}

