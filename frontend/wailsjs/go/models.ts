export namespace main {
	
	export class Host {
	    identity: string;
	    name: string;
	    info: string;
	    use: string;
	
	    static createFrom(source: any = {}) {
	        return new Host(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.identity = source["identity"];
	        this.name = source["name"];
	        this.info = source["info"];
	        this.use = source["use"];
	    }
	}

}

