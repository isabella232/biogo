// Copyright ©2011-2012 Dan Kortschak <dan.kortschak@adelaide.edu.au>
// 
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
// 
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
// 
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package alignment

import (
	"code.google.com/p/biogo/exp/alphabet"
	"code.google.com/p/biogo/exp/feat"
	"code.google.com/p/biogo/exp/seq"
	"code.google.com/p/biogo/exp/seq/protein"
	"code.google.com/p/biogo/exp/seq/sequtils"
	"fmt"
)

var (
	m, n    *Seq
	aligned = func(a *Seq) {
		for i := 0; i < a.Rows(); i++ {
			s := a.Get(i)
			fmt.Printf("%v\n", s)
		}
		fmt.Println()
		fmt.Println(a)
	}
)

func init() {
	var err error
	m, err = NewSeq("example alignment",
		[]string{"seq 1", "seq 2", "seq 3"},
		[][]alphabet.Letter{
			[]alphabet.Letter("AAA"),
			[]alphabet.Letter("CCC"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("CGA"),
			[]alphabet.Letter("TTT"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("AAA"),
			[]alphabet.Letter("CCC"),
			[]alphabet.Letter("TCG"),
			[]alphabet.Letter("TTT"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("TCC"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("CCC"),
			[]alphabet.Letter("AGT"),
			[]alphabet.Letter("CCC"),
			[]alphabet.Letter("GAA"),
			[]alphabet.Letter("TTT"),
		},
		alphabet.Protein,
		seq.DefaultConsensus)

	if err != nil {
		panic(err)
	}
}

func ExampleNewSeq() {
	m, err := NewSeq("example alignment",
		[]string{"seq 1", "seq 2", "seq 3"},
		[][]alphabet.Letter{
			[]alphabet.Letter("AAA"),
			[]alphabet.Letter("CCC"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("CGA"),
			[]alphabet.Letter("TTT"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("AAA"),
			[]alphabet.Letter("CCC"),
			[]alphabet.Letter("TCG"),
			[]alphabet.Letter("TTT"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("TCC"),
			[]alphabet.Letter("GGG"),
			[]alphabet.Letter("CCC"),
			[]alphabet.Letter("AGT"),
			[]alphabet.Letter("CCC"),
			[]alphabet.Letter("GAA"),
			[]alphabet.Letter("TTT"),
		},
		alphabet.Protein,
		seq.DefaultConsensus)
	if err != nil {
		panic(err)
	}

	aligned(m)
	// Output:
	// ACGCTGACTTGGTGCACGT
	// ACGGTGACCTGGCGCGCAT
	// ACGATGACGTGGCGCTCAT
	// 
	// acgxtgacxtggcgcxcat
}

func ExampleSeq_Add() {
	fmt.Printf("%v %v\n", m.Rows(), m)
	m.Add(protein.NewQSeq("example Protein",
		[]alphabet.QLetter{{'a', 40}, {'c', 39}, {'g', 40}, {'C', 38}, {'t', 35}, {'g', 20}},
		alphabet.Protein, alphabet.Sanger))
	fmt.Printf("%v %v\n", m.Rows(), m)
	// Output:
	// 3 acgxtgacxtggcgcxcat
	// 4 acgctgacxtggcgcxcat
}

func ExampleSeq_Copy() {
	n = m.Copy().(*Seq)
	n.Set(seq.Position{Col: 3, Row: 2}, alphabet.QLetter{L: 't'})
	aligned(m)
	fmt.Println()
	aligned(n)
	// Output:
	// ACGCTGACTTGGTGCACGT
	// ACGGTGACCTGGCGCGCAT
	// ACGATGACGTGGCGCTCAT
	// acgCtg-------------
	// 
	// acgctgacxtggcgcxcat
	// 
	// ACGCTGACTTGGTGCACGT
	// ACGGTGACCTGGCGCGCAT
	// ACGtTGACGTGGCGCTCAT
	// acgCtg-------------
	// 
	// acgctgacxtggcgcxcat
}

func ExampleSeq_Count() {
	fmt.Println(m.Rows())
	// Output:
	// 4
}

func ExampleSeq_Join() {
	aligned(n)
	err := sequtils.Join(n, m, seq.End)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	aligned(n)
	// Output:
	// ACGCTGACTTGGTGCACGT
	// ACGGTGACCTGGCGCGCAT
	// ACGtTGACGTGGCGCTCAT
	// acgCtg-------------
	// 
	// acgctgacxtggcgcxcat
	// 
	// ACGCTGACTTGGTGCACGTACGCTGACTTGGTGCACGT
	// ACGGTGACCTGGCGCGCATACGGTGACCTGGCGCGCAT
	// ACGtTGACGTGGCGCTCATACGATGACGTGGCGCTCAT
	// acgCtg-------------acgCtg-------------
	// 
	// acgctgacxtggcgcxcatacgctgacxtggcgcxcat
}

func ExampleAlignment_Len() {
	fmt.Println(m.Len())
	// Output:
	// 19
}

func ExampleSeq_Reverse() {
	aligned(m)
	fmt.Println()
	m.Reverse()
	aligned(m)
	// Output:
	// ACGCTGACTTGGTGCACGT
	// ACGGTGACCTGGCGCGCAT
	// ACGATGACGTGGCGCTCAT
	// acgCtg-------------
	// 
	// acgctgacxtggcgcxcat
	// 
	// TGCACGTGGTTCAGTCGCA
	// TACGCGCGGTCCAGTGGCA
	// TACTCGCGGTGCAGTAGCA
	// -------------gtCgca
	//
	// tacxcgcggtxcagtcgca

}

type fe struct {
	s, e int
	or   feat.Orientation
	feat.Feature
}

func (f fe) Start() int                    { return f.s }
func (f fe) End() int                      { return f.e }
func (f fe) Len() int                      { return f.e - f.s }
func (f fe) Orientation() feat.Orientation { return feat.Orientation(f.or) }

type fs []feat.Feature

func (f fs) Features() []feat.Feature { return []feat.Feature(f) }

func ExampleSeq_Stitch() {
	f := fs{
		&fe{s: -1, e: 4},
		&fe{s: 30, e: 38},
	}
	aligned(n)
	fmt.Println()
	if err := sequtils.Stitch(n, n, f); err == nil {
		aligned(n)
	} else {
		fmt.Println(err)
	}
	// Output:
	// ACGCTGACTTGGTGCACGTACGCTGACTTGGTGCACGT
	// ACGGTGACCTGGCGCGCATACGGTGACCTGGCGCGCAT
	// ACGtTGACGTGGCGCTCATACGATGACGTGGCGCTCAT
	// acgCtg-------------acgCtg-------------
	//
	// acgctgacxtggcgcxcatacgctgacxtggcgcxcat
	//
	// ACGCGTGCACGT
	// ACGGGCGCGCAT
	// ACGtGCGCTCAT
	// acgC--------
	//
	// acgcgcgcxcat
}

func ExampleSeq_Truncate() {
	aligned(m)
	err := sequtils.Truncate(m, m, 4, 12)
	if err == nil {
		fmt.Println()
		aligned(m)
	}
	// Output:
	// TGCACGTGGTTCAGTCGCA
	// TACGCGCGGTCCAGTGGCA
	// TACTCGCGGTGCAGTAGCA
	// -------------gtCgca
	//
	// tacxcgcggtxcagtcgca
	//
	// CGTGGTTC
	// CGCGGTCC
	// CGCGGTGC
	// --------
	//
	// cgcggtxc
}
