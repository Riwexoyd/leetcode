namespace csharp
{
    internal class Task2 : ILeetCodeTask
    {
        public class ListNode {
            public int val;
            public ListNode next;
            public ListNode(int val=0, ListNode next=null) 
            {
                 this.val = val;
                 this.next = next;
            }
        }

        public void Run()
        {
            ListNode l1 = new ListNode()
            {
                val = 1,
                next = new ListNode()
                {
                    val = 8,
                }
            };

            ListNode l2 = new ListNode()
            {
                val = 0
            };

            var current = AddTwoNumbers(l1, l2);

            while (current != null)
            {
                Console.WriteLine(current.val);
                current = current.next;
            }
        }

        public ListNode AddTwoNumbers(ListNode l1, ListNode l2)
        {
            int sum = 0;
            ListNode? first = null;

            ListNode? current = null;

            while (l1 != null || l2 != null || sum != 0)
            {
                if (current == null)
                {
                    first = new ListNode();
                    current = first;
                } else
                {
                    current.next = new ListNode();
                    current = current.next;
                }

                if (l1 != null)
                {
                    sum += l1.val;
                    l1 = l1.next;
                }

                if (l2 != null)
                {
                    sum += l2.val;
                    l2 = l2.next;
                }

                current.val = sum % 10;
                sum /= 10;
            }

            return first;
        }
    }
}
