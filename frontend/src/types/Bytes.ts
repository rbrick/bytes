


export class BytesPrice {

    
    priceEther: number;
    priceUSD: number;

    formattedEther: string;
    formattedUSD: string;
  
  
    constructor(priceEther: string, priceUSD: string) {
        this.priceEther = Number.parseFloat(priceEther);
        this.priceUSD = Number.parseFloat(priceUSD);

        this.formattedEther = this.priceEther.toFixed(5);
        this.formattedUSD = this.priceUSD.toFixed(2);
    }
  }
  
  
  