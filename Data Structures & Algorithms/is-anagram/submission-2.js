class Solution {
    /**
     * @param {string} s
     * @param {string} t
     * @return {boolean}
     */
    isAnagram(s, t) {
    if (s.length === t.length){const s2 = [];
    const t2 = t.split("");
    s.split("").forEach(function (val, i) {
        const find = t2.find((value) => val === value);
            if (find) {
                s2.push(find);
                const findIndex = t2.findIndex((v) => v === find);
                t2.splice(findIndex, 1);}
  });
  if (s === s2.join(''))  return true;
  else return false
} else return false
    }
}
