export namespace config {
	
	export class Config {
	    DBType: string;
	    DBPath: string;
	    NotesPath: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.DBType = source["DBType"];
	        this.DBPath = source["DBPath"];
	        this.NotesPath = source["NotesPath"];
	    }
	}

}

export namespace debt {
	
	export class Debt {
	    id: number;
	    name: string;
	    type: string;
	    total: string;
	    monthlyMin: string;
	    monthlyActual: string;
	    interest: string;
	    dueDay: number;
	
	    static createFrom(source: any = {}) {
	        return new Debt(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.total = source["total"];
	        this.monthlyMin = source["monthlyMin"];
	        this.monthlyActual = source["monthlyActual"];
	        this.interest = source["interest"];
	        this.dueDay = source["dueDay"];
	    }
	}

}

