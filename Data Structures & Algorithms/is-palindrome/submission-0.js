class Solution {
    /**
     * @param {string} s
     * @return {boolean}
     */
    isPalindrome(s) {
        const mutate = s.replace(/[^a-zA-Z0-9]/g, "").toLowerCase();
        const mutateReversed = s.replace(/[^a-zA-Z0-9]/g, "")
            .toLowerCase()
            .split("")
            .reverse()
            .join("");
        return mutate === mutateReversed ?  true :  false;
        }
        
}
