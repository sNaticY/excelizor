using System.Collections.Generic;

namespace Configs
{
    public class MultiNestedTypes 
    {
        public int Id { get; set; }
        public Dictionary<string, Dictionary<string, string>> DictTest5 { get; set; }
        public List<List<string>> ListTest3 { get; set; }
        public List<List<string>> ListTest4 { get; set; }
        public List<List<string>> ListTest5 { get; set; }
        public Dictionary<string, List<string>> DictTest6 { get; set; }
        public Dictionary<string, List<string>> DictTest7 { get; set; }
        public List<Dictionary<string, string>> ListTest6 { get; set; }
        public List<Dictionary<string, string>> ListTest7 { get; set; }
    }
}