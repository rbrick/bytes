export class Citizen {
    id: number;
    season: number;
    image: string;


    constructor(jsonObject: any) {
        this.id = jsonObject.id;
        this.season = jsonObject.season;
        this.image = jsonObject.image;
    }
}

export class Wallet {
    address: string;
    tokenBalance: string;
    totalStake: string;
    pendingRewards: string;
    stakedCitizens: Citizen[] = [];

    constructor(address: string, jsonObject: any) {
         this.address = address;
         this.tokenBalance = jsonObject.balance;
         this.totalStake = jsonObject.totalStake;
         this.pendingRewards = jsonObject.pendingRewards;

         if (jsonObject.stakedCitizens) {
            let stakedCitizens: any[] = jsonObject.stakedCitizens;
            this.stakedCitizens = stakedCitizens.map( (val) => {
               return new Citizen(val);
            });
         }
    }
}

