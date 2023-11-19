



let currencyFormatter = new Intl.NumberFormat('en-US', {currency: 'USD', style: 'currency'});

let numberFormat = new Intl.NumberFormat('en-US');

export class BytesPrice {

    
    priceEther: number;
    priceUSD: number;

    formattedEther: string;
    formattedUSD: string;


    marketCap: string;
    totalSupply: string
  
  
    constructor(priceEther: string, priceUSD: string, marketCap: string, totalSupply: string) {
        this.priceEther = Number.parseFloat(priceEther);
        this.priceUSD = Number.parseFloat(priceUSD);

        this.formattedEther = numberFormat.format(this.priceEther);
        this.formattedUSD = currencyFormatter.format(this.priceUSD);

        let mc = Number.parseFloat(marketCap);
        let ts = Number.parseFloat(totalSupply);

        this.marketCap = currencyFormatter.format(mc);
        this.totalSupply = numberFormat.format(ts);
    }
  }
  
  
  