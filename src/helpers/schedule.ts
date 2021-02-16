export const convertPeriodsFromStringToArray = (periodString: string): number[] => {
  const PERIOD_REGEX = new RegExp("^(\\d{1,2})-(\\d{1,2})$");
  if (PERIOD_REGEX.test(periodString)) {
    periodString.match(PERIOD_REGEX);
    const first = Math.min(Number(RegExp.$1), Number(RegExp.$2));
    const last = Math.max(Number(RegExp.$1), Number(RegExp.$2));
    const periods = [];
    for (let i = first; i <= last; i++) {
      periods.push(i);
    }
    return periods;
  }
  return [];
};
