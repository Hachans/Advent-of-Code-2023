class Program
{
    static void Main(string[] args)
    {
        List<string> lines = readInput("../input.txt");
        Dictionary<string, char> myDict = initDict();
        int total = 0;

        foreach (var line in lines)
        {
            char first = extractFirst(line, myDict);
            char last = extractLast(line, myDict);
            string s = string.Concat(first, last);
            int sum = int.Parse(s);
            total += sum;
        }
        Console.WriteLine($"Total sum: {total}");
    }

    static List<string> readInput(string path)
    {
        List<string> lineList = new List<string>();
        var file = File.ReadLines(path);
        foreach (var line in file)
        {
            lineList.Add(line);
        }
        return lineList;
    }

    static Dictionary<string, char> initDict()
    {

        Dictionary<string, char> myDict = new Dictionary<string, char>();

        myDict.Add("one", '1');
        myDict.Add("two", '2');
        myDict.Add("three", '3');
        myDict.Add("four", '4');
        myDict.Add("five", '5');
        myDict.Add("six", '6');
        myDict.Add("seven", '7');
        myDict.Add("eight", '8');
        myDict.Add("nine", '9');

        return myDict;
    }

    static char extractFirst(string line, Dictionary<string, char> myDict)
    {
        for (int i = 0; i < line.Length; i++)
        {
            if (Char.IsDigit(line[i]))
            {
                return line[i];
            }
            else
            {
                foreach (KeyValuePair<string, char> entry in myDict)
                {
                    if (line.Substring(0, i + 1).Contains(entry.Key))
                    {
                        return entry.Value;
                    }
                }
            }
        }
        return ' ';
    }

    static char extractLast(string line, Dictionary<string, char> myDict)
    {
        for (int i = line.Length - 1; i >= 0; i--)
        {
            if (Char.IsDigit(line[i]))
            {
                return line[i];
            }
            else
            {
                foreach (KeyValuePair<string, char> entry in myDict)
                {
                    if (line.Substring(i, line.Length - i).Contains(entry.Key))
                    {
                        return entry.Value;
                    }
                }
            }
        }
        return ' ';
    }
}
