using System.Collections.Generic;

namespace Configs
{
    public class MultiNestedTypes 
    {
        public int Id;
        public Dictionary<string, Dictionary<string, int>> DictTest3;
        public Dictionary<string, Dictionary<string, string>> DictTest5;
        public List<List<string>> ListTest3;
        public List<List<string>> ListTest4;
        public List<List<string>> ListTest5;
        public Dictionary<string, List<string>> DictTest6;
        public Dictionary<string, List<string>> DictTest7;
        public List<Dictionary<string, string>> ListTest6;
        public List<Dictionary<string, string>> ListTest7;
    }
}