class Solution {
    public int[] solution(int numer1, int denom1, int numer2, int denom2) {
        
        int rden,rnum;
        if (denom2%denom1 == 0){
            rden = denom2;
            rnum = (numer1 * (denom2/denom1)) + numer2;
        }else {
            rden = denom1*denom2;
            rnum = (numer1 * denom2) + (numer2*denom1);
        }
        
        for(int i=2; i<=rden; i++){
            if (rnum%i==0 && rden%i==0){
                rnum = rnum/i;
                rden = rden/i;
                i=1;
            }
        }
        
        int[] answer = {rnum,rden};
        return answer;
        
    }
}